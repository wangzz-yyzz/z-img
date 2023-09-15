package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"z-img/api"
	"z-img/config"
)

// InitRouter NewRouter creates a new router
func InitRouter() *gin.Engine {
	// set running mode and create a router
	gin.SetMode(config.RunMode)
	r := gin.New()

	// load middleware
	r.Use(cors.Default())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// load tmpl file
	r.LoadHTMLGlob("tmpl/index.tmpl")

	// load img api
	imgApi := r.Group("/img")
	imgApi.GET("/get", api.GetImgById)
	imgApi.POST("/post", api.PostImg)
	imgApi.GET("/delete", api.RemoveImgById)

	// load ui api
	uiApi := r.Group("/ui")
	uiApi.GET("/index", api.GetIndex)
	uiApi.GET("/get_list", api.GetMd5List)
	uiApi.StaticFS("/static", http.Dir("tmpl/static"))

	// load static file server
	r.StaticFS("/list", http.Dir(config.ImgPath))

	return r
}
