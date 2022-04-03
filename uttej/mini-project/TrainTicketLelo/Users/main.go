package main

import (
	"Users/config"
	controller "Users/controllers"
	"Users/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/signup", controller.CreateUser)
	r.POST("/users/login", controller.LogInUser)
	r.GET("/users", middleware.IfAuthorized(controller.GetAllUsers))
	r.GET("/users/:userId", middleware.IfAuthorized(controller.GetUserById))
	r.PUT("/users/:userId", middleware.IfAuthorized(controller.UpdateUserById))
	r.DELETE("/users/:userId", middleware.IfAuthorized(controller.DeleteUserbyId))

	portAddress := ":" + config.UserServicePort
	r.Run(portAddress)
}
