package routes

import (
	"tejas/controllers"
	"tejas/middlewares"

	"github.com/gin-gonic/gin"
)

func ScheduleRoutes(router *gin.Engine) {
	public := router.Group("/api/schedule")
	public.GET("/", controllers.Availabilty())

	private := router.Group("/api/schedule")
	private.Use(middlewares.AuthenticateJWT())
	private.POST("reserve", controllers.ReserveSeat())

	adminOnly := router.Group("/api/schedule")
	adminOnly.Use(middlewares.AuthenticateJWT())
	adminOnly.Use(middlewares.OnlyAdmin())
	adminOnly.POST("/add", controllers.AddTrainSchedule())
}
