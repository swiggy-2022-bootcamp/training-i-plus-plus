package routes

import (
	"github.com/gin-gonic/gin"
	"sanitaria/controllers"
	"sanitaria/middlewares"
)

func PatientRoutes(router *gin.Engine){	
	router.POST("/patientRegistration",controllers.RegisterPatient())
	router.POST("/patientLogin",controllers.LoginPatient())
	router.Use(middlewares.AuthenticateJWT())
	router.GET("/patient/:id",controllers.GetPatientByID())
	router.PUT("/patient/:id", controllers.EditPatientByID())
	router.DELETE("/patient/:id", controllers.DeletePatientByID())
	router.GET("/patients", controllers.GetAllPatients())
}
