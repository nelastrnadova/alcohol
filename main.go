package main

import (
	"alcohol/controller"
	"alcohol/service"
	"github.com/gin-gonic/gin"
)

var (
	postService service.PostService = service.New()
	postController controller.PostController = controller.New(postService)
)

func main() {
	r := gin.Default()

	r.GET("/posts", func(ctx *gin.Context) {
		ctx.JSON(200, postController.FindAll())
	})

	r.POST("/posts", func(ctx *gin.Context) {
		ctx.JSON(200, postController.Save(ctx))
	})

	r.Run()
}