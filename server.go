package main

import (
	"hello/middlewares"
	"hello/controller"
	"hello/service"
	"github.com/gin-gonic/gin"
)

var (
	videoService service.VideoService = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.New()
	server.Use(gin.Recovery())
	server.Use(middlewares.Logger())
	// server.Use(gin.Logger())

	// server.GET("/test", func (ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"message": "OK",
	// 	})
	// })

	server.GET("/posts", func (ctx *gin.Context){
		ctx.JSON(200, videoController.FindAll())
	}) 

	server.POST("/posts", func (ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.Run(":8080")
}