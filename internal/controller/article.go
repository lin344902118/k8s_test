package controller

import (
	"github.com/gin-gonic/gin"
	"go_api_framework/internal/service"
	"go_api_framework/pkg/errcode"
	"net/http"
)

type ArticleController struct {
}

func NewArticleController() *ArticleController {
	return &ArticleController{}
}

func (ac ArticleController) ArticleList(c *gin.Context) {
	svc := service.New(c.Request.Context())
	articles, err := svc.GetArticleList(1, 20)
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

func (ac ArticleController) GetArticle(c *gin.Context) {
	svc := service.New(c.Request.Context())
	aId, exist := c.Get("id")
	if !exist {
		c.HTML(http.StatusOK, "layout.html", gin.H{
			"error": errcode.MissingArticleId.Msg(),
		})
	}

	article, err := svc.GetArticle(aId.(uint32))
	if err != nil {
		c.HTML(http.StatusOK, "layout.html", gin.H{
			"error": errcode.GetArticleError.Msg(),
		})
	} else {
		c.HTML(http.StatusOK, "articleDetail.html", gin.H{
			"blog": article,
		})
	}
}

func (ac ArticleController) UpdateArticle(c *gin.Context) {
}

func (ac ArticleController) CreateArticle(c *gin.Context) {

}

func (ac ArticleController) DeleteArticle(c *gin.Context) {
	svc := service.New(c.Request.Context())
	aId, exist := c.Get("id")
	if !exist {
		c.HTML(http.StatusOK, "layout.html", gin.H{
			"error": errcode.MissingArticleId.Msg(),
		})
	}
	err := svc.DeleteArticle(aId.(uint32))
	if err != nil {
		c.HTML(http.StatusOK, "layout.html", gin.H{
			"error": errcode.GetArticleError.Msg(),
		})
	} else {
		c.Redirect(http.StatusOK, "/article")
	}
}
