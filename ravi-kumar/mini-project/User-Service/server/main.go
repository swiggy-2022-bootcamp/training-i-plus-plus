package main

import (
	"User-Service/config"
	"User-Service/controller"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/users", controller.CreateUser)
	r.GET("/users", controller.GetAllUsers)
	r.GET("/users/:userId", controller.GetUserById)
	r.PUT("/users/:userId", controller.UpdateUserById)
	r.DELETE("/users/:userId", controller.DeleteUserbyId)

	portAddress := ":" + strconv.Itoa(config.USER_SERVICE_SERVER_PORT)
	r.Run(portAddress)
}
