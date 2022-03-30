package main

import (
	"User-Service/config"
	"User-Service/controller"
	"User-Service/middleware"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/users", controller.CreateUser)
	r.POST("/users/login", controller.LogInUser)
	r.GET("/users", middleware.IfAuthorized(controller.GetAllUsers))
	r.GET("/users/:userId", middleware.IfAuthorized(controller.GetUserById))
	r.PUT("/users/:userId", middleware.IfAuthorized(controller.UpdateUserById))
	r.DELETE("/users/:userId", middleware.IfAuthorized(controller.DeleteUserbyId))

	portAddress := ":" + strconv.Itoa(config.USER_SERVICE_SERVER_PORT)
	r.Run(portAddress)
}
