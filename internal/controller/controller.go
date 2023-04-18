package controller

import (
	"github.com/gin-gonic/gin"
	"go_api_framework/internal/service"
	"net/http"
)

type MainController struct {
}

func NewMainController() *MainController {
	return &MainController{}
}

func (m *MainController) Index(c *gin.Context) {
	svc := service.New(c.Request.Context())
	articles, err := svc.GetAllArticleWithCategorys(1, 20)
	if err != nil {
		c.HTML(http.StatusOK, "home.html", gin.H{
			"error": "获取文章失败",
		})
	} else {
		c.HTML(http.StatusOK, "home.html", gin.H{
			"blogs": articles,
		})
	}

}

func (m *MainController) About(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", gin.H{})
}

func (m *MainController) Search(c *gin.Context) {
	key := c.Query("q")
	svc := service.New(c.Request.Context())
	datas, err := svc.Search(key)
	if err != nil {
		c.HTML(http.StatusOK, "home.html", gin.H{
			"error": "搜索文章失败",
		})
	} else {
		c.HTML(http.StatusOK, "home.html", gin.H{
			"blogs": datas,
		})
	}
}
