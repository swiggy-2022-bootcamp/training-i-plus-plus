package routes

import (
	"tejas/controllers"
	"tejas/middlewares"

	"github.com/gin-gonic/gin"
)

func TrainRoutes(router *gin.Engine) {
	adminOnly := router.Group("/api/train")
	adminOnly.Use(middlewares.AuthenticateJWT())
	adminOnly.Use(middlewares.OnlyAdmin())
	adminOnly.POST("/add", controllers.AddTrain())
	adminOnly.DELETE("/remove", controllers.RemoveTrain())
}
