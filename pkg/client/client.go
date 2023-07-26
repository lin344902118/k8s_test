package client

import (
	"k8s.io/client-go/kubernetes"
	"k8s_test/global"
)

type Clients struct {
	clientSet kubernetes.Interface
}

// 根据全局kubeconfig创建一个新的ClientSet
func NewClients() (*Clients) {
	clientSet, err := kubernetes.NewForConfig(global.KubeConfig)
	if err != nil {
		return nil
	}
	return &Clients{
		clientSet: clientSet,
	}
}

func (c *Clients) ClientSet() kubernetes.Interface {
	return c.clientSet
}
