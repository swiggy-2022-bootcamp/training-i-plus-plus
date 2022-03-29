package routes

import (
	"tejas/controllers"
	"tejas/middlewares"

	"github.com/gin-gonic/gin"
)

func TrainRoutes(router *gin.Engine) {
	private := router.Group("/api/train")
	private.Use(middlewares.AuthenticateJWT())
	private.Use(middlewares.OnlyAdmin())
	private.POST("/add", controllers.AddTrain())
	private.POST("/remove", controllers.RemoveTrain())
}
