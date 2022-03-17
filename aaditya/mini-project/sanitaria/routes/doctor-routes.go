package routes

import (
	"github.com/gin-gonic/gin"
	"sanitaria/controllers"
)

func DoctorRoutes(router *gin.Engine){	
	router.POST("/doctorRegistration",controllers.RegisterDoctor())
	router.GET("/doctor/:id",controllers.GetDoctorByID())
	router.PUT("/doctor/:id", controllers.EditDoctorByID())
	router.DELETE("/doctor/:id", controllers.DeleteDoctorByID())
	router.GET("/doctors", controllers.GetAllDoctors())
}
