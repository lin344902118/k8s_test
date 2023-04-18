package routers

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"go_api_framework/internal/controller"
	"path/filepath"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.HTMLRender = loadTemplate("internal/html")
	r.Static("/static", "internal/static")

	c := controller.NewMainController()
	r.GET("/", c.Index)
	r.GET("/index", c.Index)
	r.GET("/about", c.About)
	r.GET("/search", c.Search)
	ac := controller.NewArticleController()
	r.GET("/article", ac.ArticleList)
	cc := controller.NewCategoryController()
	r.GET("/category", cc.CategoryList)
	return r
}

func loadTemplate(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	layout, err := filepath.Glob(templatesDir + "/layout.html")
	if err != nil {
		panic(err.Error())
	}
	contents, err := filepath.Glob(templatesDir + "/content/*.html")
	if err != nil {
		panic(err.Error())
	}
	// Generate our templates map from our layouts/ and articles/ directories
	for _, include := range contents {
		files := append(layout, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}
