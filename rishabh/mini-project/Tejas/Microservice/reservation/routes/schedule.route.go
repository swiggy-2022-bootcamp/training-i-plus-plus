package routes

import (
	"reservationService/controllers"
	"reservationService/middlewares"

	"github.com/gin-gonic/gin"
)

func ScheduleRoutes(router *gin.Engine) {
	public := router.Group("/api/schedule")
	public.Use(middlewares.LoggerMiddleware("Schedule"))

	public.GET("/", controllers.Availabilty())

	private := router.Group("/api/schedule")
	private.Use(middlewares.AuthenticateJWT())
	private.Use(middlewares.LoggerMiddleware("Schedule"))

	private.POST("reserve", controllers.ReserveSeat())

	adminOnly := router.Group("/api/schedule")
	adminOnly.Use(middlewares.AuthenticateJWT())
	adminOnly.Use(middlewares.OnlyAdmin())
	adminOnly.Use(middlewares.LoggerMiddleware("Schedule"))

	adminOnly.POST("/add", controllers.AddTrainSchedule())
}
