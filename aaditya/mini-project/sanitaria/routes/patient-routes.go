package routes

import (
	"github.com/gin-gonic/gin"
	"sanitaria/controllers"
)

func PatientRoutes(router *gin.Engine){	
	router.POST("/patientRegistration",controllers.RegisterPatient())
	router.GET("/patient/:id",controllers.GetPatientByID())
	router.PUT("/patient/:id", controllers.EditPatientByID())
	router.DELETE("/patient/:id", controllers.DeletePatientByID())
	router.GET("/patients", controllers.GetAllPatients())
}
