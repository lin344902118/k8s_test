package informer

import (
	"k8s.io/client-go/informers"
	"k8s_test/pkg/client"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"time"
)

var sharedInformerFactory informers.SharedInformerFactory

func NewSharedInformerFactory(stop <-chan struct{}) (err error) {
	// 加载客户端
	clients := client.NewClients()

	options := informers.WithNamespace("kube-system")
	// 实例化
	sharedInformerFactory = informers.NewSharedInformerFactoryWithOptions(clients.ClientSet(), 60 * time.Second, options)

	// 启动informer
	gvrs := []schema.GroupVersionResource{
		{Group:"", Version:"v1", Resource: "pods"},
		{Group:"", Version:"v1", Resource: "services"},
		{Group:"", Version:"v1", Resource: "namespaces"},
	}

	for _, v := range gvrs {
		// 创建informer
		_, err = sharedInformerFactory.ForResource(v)
		if err != nil {
			return
		}
	}
	sharedInformerFactory.Start(stop)
	sharedInformerFactory.WaitForCacheSync(stop)

	return
}

func Get() informers.SharedInformerFactory {
	return sharedInformerFactory
}
