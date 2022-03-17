package routes

import (
	"github.com/gin-gonic/gin"
	"sanitaria/controllers"
)

func GeneralUserRoutes(router *gin.Engine){	
	router.POST("/generalUserRegistration",controllers.RegisterGeneralUser())
	router.GET("/generalUser/:id",controllers.GetGeneralUserByID())
	router.PUT("/generalUser/:id", controllers.EditGeneralUserByID())
	router.DELETE("/generalUser/:id", controllers.DeleteGeneralUserByID())
	router.GET("/generalUsers", controllers.GetAllGeneralUsers())
}
