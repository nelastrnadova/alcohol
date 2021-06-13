package main

import (
	"alcohol/controller"
	"alcohol/middlewares"
	"alcohol/service"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
	"io"
	"os"
)

var (
	postService service.PostService = service.New()
	postController controller.PostController = controller.New(postService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()

	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())

	server.GET("/posts", func(ctx *gin.Context) {
		ctx.JSON(200, postController.FindAll())
	})

	server.POST("/posts", func(ctx *gin.Context) {
		ctx.JSON(200, postController.Save(ctx))
	})

	server.Run()
}