package routes

import (
	"ticket_reservation_system/controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/user", controller.AddUser())
	router.GET("/user/:username", controller.GetUserByUsername())
	router.PUT("/user/:username", controller.UpdateUserPassword())
	router.DELETE("/user/:username", controller.DeleteUserByUsername())

	router.GET("/users", controller.GetUsers())

}
