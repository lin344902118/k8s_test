package controller

import (
	"github.com/gin-gonic/gin"
	"go_api_framework/internal/service"
	"net/http"
)

type CategoryController struct {
}

func NewCategoryController() *CategoryController {
	return &CategoryController{}
}

func (cc CategoryController) CategoryList(c *gin.Context) {
	svc := service.New(c.Request.Context())
	categorys, err := svc.GetCategoryList(1, 20)
	if err != nil {
		c.HTML(http.StatusOK, "categoryList.html", gin.H{
			"error": "获取分类失败",
		})
	} else {
		c.HTML(http.StatusOK, "categoryList.html", gin.H{
			"categorys": categorys,
		})
	}
}
