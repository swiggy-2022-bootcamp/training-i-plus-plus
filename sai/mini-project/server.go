package main

import (
	"io"
	"net/http"
	"os"

	"cherie.com/golang-gin/controller"
	"cherie.com/golang-gin/middlewares"
	"cherie.com/golang-gin/service"
	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() { // used to log output to a file on disk
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	// server := gin.Default()

	setupLogOutput()

	server := gin.New()
	server.Use(gin.Recovery())
	server.Use(middlewares.Logger()) //, middlewares.BasicAuth())
	// server.Use(gindump.Dump())

	// server.GET("/test", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"message": "OK",
	// 	})
	// })

	apiRoute := server.Group("/api")
	{

		apiRoute.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})

		apiRoute.POST("/video", middlewares.BasicAuth(), func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video Input is Valid"})
			}
		})
	}

	server.Run(":8080")
}
