package global

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s_test/pkg/logger"
	"k8s_test/pkg/setting"
)

var (
	ServerSetting    *setting.ServerSetting
	AppSetting       *setting.AppSetting
	DatabaseSetting  *setting.DataBaseSetting
	KuberneteSetting *setting.KubeneteSetting

	Logger     *logger.Logger
	KubeConfig *rest.Config
	ClientSet  *kubernetes.Clientset
)
