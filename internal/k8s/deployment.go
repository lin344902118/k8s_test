package k8s

import (
	"context"
	"fmt"
	"k8s_test/global"
	"k8s_test/pkg/client"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/util/retry"
)

type ContainerDemo struct {
	Name  string
	Image string
	Port  int32
}

type DeploymentTemplate struct {
	Labels    map[string]string
	Container ContainerDemo
}

type DeploymentDemo struct {
	Name        string
	Replicas    int32
	MatchLabels map[string]string
	Template    DeploymentTemplate
}

func NewDeployment(name, imageName, imageAddr string, replicas, port int32) DeploymentDemo {
	deploy := DeploymentDemo{
		Name:        name,
		Replicas:    replicas,
		MatchLabels: map[string]string{"app": name},
		Template: DeploymentTemplate{
			Labels: map[string]string{"app": name},
			Container: ContainerDemo{
				Name:  imageName,
				Image: imageAddr,
				Port:  port,
			},
		},
	}
	return deploy
}

func CreateDeployment(ctx context.Context, demo DeploymentDemo) error {
	global.Logger.Infof(ctx, "demo:%+v", demo)
	deploymentsClient := client.NewClients().ClientSet().AppsV1().Deployments(apiv1.NamespaceDefault)
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: demo.Name,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(demo.Replicas),
			Selector: &metav1.LabelSelector{
				MatchLabels: demo.MatchLabels,
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: demo.Template.Labels,
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  demo.Template.Container.Name,
							Image: demo.Template.Container.Image,
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: demo.Template.Container.Port,
								},
							},
						},
					},
				},
			},
		},
	}
	// Create Deployment
	global.Logger.Infof(ctx, "Creating deployment...")
	result, err := deploymentsClient.Create(ctx, deployment, metav1.CreateOptions{})
	if err != nil {
		global.Logger.Errorf(ctx, "deployments client create deployment failed.err:%s", err.Error())
		return err
	}
	global.Logger.Infof(ctx, "Created deployment %q.\n", result.GetObjectMeta().GetName())
	return nil
}

func ListDeployment(ctx context.Context) ([]DeploymentDemo, error) {
	deploymentsClient := client.NewClients().ClientSet().AppsV1().Deployments(apiv1.NamespaceDefault)
	list, err := deploymentsClient.List(ctx, metav1.ListOptions{})
	if err != nil {
		global.Logger.Errorf(ctx, "deployments client list failed.err:%v", err)
		return []DeploymentDemo{}, err
	}
	demos := make([]DeploymentDemo, 0, len(list.Items))
	for _, d := range list.Items {
		demo := DeploymentDemo{
			Name:        d.Name,
			MatchLabels: d.Spec.Selector.MatchLabels,
			Replicas:    *d.Spec.Replicas,
			Template: DeploymentTemplate{
				Labels: d.Spec.Template.Labels,
				Container: ContainerDemo{
					Name:  d.Spec.Template.Spec.Containers[0].Name,
					Image: d.Spec.Template.Spec.Containers[0].Image,
					Port:  d.Spec.Template.Spec.Containers[0].Ports[0].ContainerPort,
				},
			},
		}
		demos = append(demos, demo)
	}
	return demos, nil
}

func UpdateDeployment(ctx context.Context, image string) error {
	deploymentsClient := client.NewClients().ClientSet().AppsV1().Deployments(apiv1.NamespaceDefault)
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// Retrieve the latest version of Deployment before attempting update
		// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
		result, getErr := deploymentsClient.Get(ctx, "demo-deployment", metav1.GetOptions{})
		if getErr != nil {
			return fmt.Errorf("Failed to get latest version of Deployment: %v", getErr)
		}
		// reduce replica count
		result.Spec.Template.Spec.Containers[0].Image = image // change nginx version
		_, updateErr := deploymentsClient.Update(ctx, result, metav1.UpdateOptions{})
		return updateErr
	})
	if retryErr != nil {
		return fmt.Errorf("Update deployment failed. err:%v", retryErr)
	}
	return nil
}

func DeleteDeployment(ctx context.Context, name string) error {
	global.Logger.Info(ctx, "Deleting deployment...")
	deletePolicy := metav1.DeletePropagationForeground
	deploymentsClient := client.NewClients().ClientSet().AppsV1().Deployments(apiv1.NamespaceDefault)
	if err := deploymentsClient.Delete(ctx, name, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		return fmt.Errorf("Delete deployment failed.err: %v", err)
	}
	global.Logger.Info(ctx, "Deleted deployment.")
	return nil
}

func int32Ptr(i int32) *int32 { return &i }
