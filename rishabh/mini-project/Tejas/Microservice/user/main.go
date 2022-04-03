package main

import (
	"userService/configs"
	"userService/routes"
	"userService/services"

	"github.com/gin-gonic/gin"
)

var logger = services.NewLoggerService("User Service")

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server is up and running",
		})
	})

	// init db
	configs.ConnectDB()
	logger.Log("Connected to DB")

	// init routes
	routes.UserRoutes(router)

	router.Run("localhost:" + configs.EnvPORT())
}
