package routes

import (
	"github.com/gin-gonic/gin"
	"sanitaria-microservices/generalUserModule/controllers"
	"sanitaria-microservices/generalUserModule/middlewares"
)

func GeneralUserRoutes(router *gin.Engine){	
	public := router.Group("")
	public.POST("/generalUserRegistration",controllers.RegisterGeneralUser())
	public.POST("/generalUserLogin",controllers.LoginGeneralUser())
	private := router.Group("")
	private.Use(middlewares.AuthenticateJWT())
	private.GET("/generalUser/:id",controllers.GetGeneralUserByID())
	private.PUT("/generalUser/:id", controllers.EditGeneralUserByID())
	private.DELETE("/generalUser/:id", controllers.DeleteGeneralUserByID())
	private.GET("/generalUsers", controllers.GetAllGeneralUsers())
	// private.GET("/generalUsers/available-appointments", controllers.GetAvailableAppointments())
	// private.POST("/generalUsers/book-appointment/:id",controllers.BookAppointment())
}
