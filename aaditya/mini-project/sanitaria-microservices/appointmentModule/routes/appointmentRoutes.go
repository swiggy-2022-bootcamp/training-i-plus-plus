package routes

import (
	"github.com/gin-gonic/gin"
	"sanitaria-microservices/appointmentModule/controllers"
	"sanitaria-microservices/appointmentModule/middlewares"
)

func AppointmentRoutes(router *gin.Engine){
	private := router.Group("")
	private.Use(middlewares.AuthenticateJWT())
	private.GET("/appointments",controllers.GetAllAppointments())
	private.POST("/book-appointment/:userId",controllers.BookAppointment())
}