package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"k8s_test/internal/k8s/informer"
	"k8s.io/apimachinery/pkg/labels"
	"k8s_test/global"
	"fmt"
)

type mainController struct{}

func NewMainController() *mainController {
	return &mainController{}
}

func (controller *mainController) GetIndex(c *gin.Context) {
	stopCh := make(chan struct{})
	err := informer.NewSharedInformerFactory(stopCh)
	if err != nil {
		fmt.Printf("New Factory err:%+v\n", err)
		global.Logger.Errorf(c, "NewSharedInformerFactory failed:%+v", err)
	}
	items, err := informer.Get().Core().V1().Pods().Lister().List(labels.Everything())
	if err != nil {
		fmt.Printf("get pods err:%+v\n", err)
		global.Logger.Errorf(c,"get kube-system pods failed.err:%+v", err)
	}
	fmt.Println("items", items)
	for _, v := range items {
		fmt.Printf("namespace:%v, name:%v\n", v.Namespace, v.Name)
		global.Logger.Infof(c,"namespace:%v, name:%v\n", v.Namespace, v.Name)
	}
	c.HTML(http.StatusOK, "deployment.html", gin.H{})
}
