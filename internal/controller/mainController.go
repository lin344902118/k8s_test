package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type mainController struct{}

func NewMainController() *mainController {
	return &mainController{}
}

func (controller *mainController) GetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "deployment.html", gin.H{})
}
