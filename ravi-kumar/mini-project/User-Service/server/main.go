package main

import (
	"User-Service/config"
	"User-Service/controller"
	"User-Service/middleware"
	"User-Service/service"
	"io"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	//logs
	userSerivceLogFile, _ := os.Create("userService.log")
	gin.DefaultWriter = io.MultiWriter(os.Stdout, userSerivceLogFile)
	service.GetDefaultWriter = &gin.DefaultWriter
	service.InitLoggerService()

	r := gin.Default()
	r.Use(cors.AllowAll())

	r.POST("/users", controller.CreateUser)
	r.POST("/users/login", controller.LogInUser)
	r.GET("/users", controller.GetAllUsers)
	r.GET("/users/:userId", middleware.IfAuthorized(controller.GetUserById))
	r.PUT("/users/:userId", middleware.IfAuthorized(controller.UpdateUserById))
	r.DELETE("/users/:userId", middleware.IfAuthorized(controller.DeleteUserbyId))
	r.Static("/swaggerui", config.SWAGGER_PATH)
	portAddress := ":" + strconv.Itoa(config.USER_SERVICE_SERVER_PORT)
	r.Run(portAddress)
}
