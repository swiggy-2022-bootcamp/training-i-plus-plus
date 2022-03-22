package main

import (
	"io"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/practice/controller"
	"github.com/practice/middlewares"
	"github.com/practice/service"
)

var (
	videoService service.VideoService=service.New()
	videoController controller.VideoController=controller.New(videoService)
)

func setupLogOutput(){
	f,_:=os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f,os.Stdout)
}

func main(){
	// server:=gin.DefaNult()
	server:=gin.New()
	server.Use(gin.Recovery(),middlewares.Logger(),middlewares.BasicAuth())
	server.GET("/posts", func(ctx *gin.Context){
		ctx.JSON(200,videoController.FindAll())
	})

	server.POST("/posts",func(ctx *gin.Context){
		result,err:=videoController.Save(ctx)
		if err==nil{
			ctx.JSON(http.StatusBadRequest,result)
		} else{
			ctx.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		}
	})
	server.Run(":8080")
}