package routes

import (
	"github.com/gin-gonic/gin"
	"sanitaria/controllers"
	"sanitaria/middlewares"
)

func DoctorRoutes(router *gin.Engine){	
	router.POST("/doctorRegistration",controllers.RegisterDoctor())
	router.POST("/doctorLogin",controllers.LoginDoctor())
	router.Use(middlewares.AuthenticateJWT())
	router.GET("/doctor/:id",controllers.GetDoctorByID())
	router.PUT("/doctor/:id", controllers.EditDoctorByID())
	router.DELETE("/doctor/:id", controllers.DeleteDoctorByID())
	router.GET("/doctors", controllers.GetAllDoctors())
}
