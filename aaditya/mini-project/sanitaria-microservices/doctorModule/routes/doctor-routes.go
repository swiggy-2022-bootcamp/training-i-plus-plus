package routes

import (
	"github.com/gin-gonic/gin"
	"sanitaria-microservices/doctorModule/controllers"
	"sanitaria-microservices/doctorModule/middlewares"
)

func DoctorRoutes(router *gin.Engine){
	public := router.Group("")	
	public.POST("/doctorRegistration",controllers.RegisterDoctor())
	public.POST("/doctorLogin",controllers.LoginDoctor())
	private := router.Group("")
	private.Use(middlewares.AuthenticateJWT())
	private.GET("/doctor/:id",controllers.GetDoctorByID())
	private.PUT("/doctor/:id", controllers.EditDoctorByID())
	private.DELETE("/doctor/:id", controllers.DeleteDoctorByID())
	private.GET("/doctors", controllers.GetAllDoctors())
	private.POST("/doctors/add-appointment/:doctorId",controllers.OpenSlotsForAppointments())
}
