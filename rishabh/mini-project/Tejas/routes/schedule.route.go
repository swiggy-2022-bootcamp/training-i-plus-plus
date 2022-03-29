package routes

import (
	"tejas/controllers"
	"tejas/middlewares"

	"github.com/gin-gonic/gin"
)

func ScheduleRoutes(router *gin.Engine) {
	adminOnly := router.Group("/api/schedule")
	adminOnly.Use(middlewares.AuthenticateJWT())
	adminOnly.Use(middlewares.OnlyAdmin())
	adminOnly.POST("/add", controllers.AddTrainSchedule())
}
