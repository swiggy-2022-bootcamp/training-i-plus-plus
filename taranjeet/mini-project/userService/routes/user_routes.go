package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/taran1515/crud/controllers"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/user", controllers.CreateUser())
	router.GET("/user/:userId", controllers.GetAUser())
	router.PUT("/user/:userId", controllers.EditAUser())
	router.DELETE("/user/:userId", controllers.DeleteAUser())
	router.GET("/users", controllers.GetAllUsers())

}
