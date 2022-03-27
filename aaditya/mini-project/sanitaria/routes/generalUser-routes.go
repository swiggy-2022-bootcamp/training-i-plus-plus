package routes

import (
	"github.com/gin-gonic/gin"
	"sanitaria/controllers"
	"sanitaria/middlewares"
)

func GeneralUserRoutes(router *gin.Engine){	
	router.POST("/generalUserRegistration",controllers.RegisterGeneralUser())
	router.POST("/generalUserLogin",controllers.LoginGeneralUser())
	router.Use(middlewares.AuthenticateJWT())
	router.GET("/generalUser/:id",controllers.GetGeneralUserByID())
	router.PUT("/generalUser/:id", controllers.EditGeneralUserByID())
	router.DELETE("/generalUser/:id", controllers.DeleteGeneralUserByID())
	router.GET("/generalUsers", controllers.GetAllGeneralUsers())
	router.POST("/generalUsers/book-appointment/:id",controllers.BookAppointment())
}
