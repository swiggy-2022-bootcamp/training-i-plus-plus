package routes

import (
	"github.com/gin-gonic/gin"
	"sanitaria-microservices/patientModule/controllers"
	"sanitaria-microservices/patientModule/middlewares"
)

func PatientRoutes(router *gin.Engine){	
	public := router.Group("")
	public.POST("/patientRegistration",controllers.RegisterPatient())
	public.POST("/patientLogin",controllers.LoginPatient())
	private := router.Group("")
	private.Use(middlewares.AuthenticateJWT())
	private.GET("/patient/:id",controllers.GetPatientByID())
	private.PUT("/patient/:id", controllers.EditPatientByID())
	private.DELETE("/patient/:id", controllers.DeletePatientByID())
	private.GET("/patients", controllers.GetAllPatients())
}
