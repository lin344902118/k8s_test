package controller

import (
	"fmt"
	"k8s_test/internal/k8s"
	"net/http"
	"strconv"

	"k8s_test/global"

	"github.com/gin-gonic/gin"
)

type deploymentController struct {
}

func NewDeploymentController() *deploymentController {
	return &deploymentController{}
}

func (controller *deploymentController) ListDeployment(c *gin.Context) {
	deployments, err := k8s.ListDeployment(c)
	global.Logger.Infof(c, "deployments:%+v", deployments)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": fmt.Sprintf("list deployment failed.err:%v", err),
		})
	} else {
		c.HTML(http.StatusOK, "listDeployment.html", gin.H{
			"deployments": deployments,
		})
	}
}

func (controller *deploymentController) CreateDeployment(c *gin.Context) {
	name := c.PostForm("name")
	imageName := c.PostForm("imageName")
	imageAddr := c.PostForm("imageAddress")
	imagePort := c.PostForm("imagePort")
	replicas := c.PostForm("replicas")
	port, _ := strconv.Atoi(imagePort)
	rep, _ := strconv.Atoi(replicas)
	deploy := k8s.NewDeployment(name, imageName, imageAddr, int32(rep), int32(port))
	global.Logger.Infof(c, "deployment:%+v", deploy)
	err := k8s.CreateDeployment(c, deploy)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": fmt.Sprintf("create deployment failed.err:%v", err),
		})
	} else {
		c.Redirect(http.StatusFound, "/deployment")
	}

}

func (controller *deploymentController) DeleteDeployment(c *gin.Context) {
	name := c.Query("name")
	err := k8s.DeleteDeployment(c, name)
	if err != nil {
		c.JSONP(http.StatusInternalServerError, err)
	} else {
		c.JSONP(http.StatusOK, "success")
	}
}
