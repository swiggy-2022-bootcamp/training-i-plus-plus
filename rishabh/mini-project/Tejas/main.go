package main

import (
	"tejas/configs"
	"tejas/routes"
	"tejas/services"

	"github.com/gin-gonic/gin"
)

var logger = services.NewLoggerService("main")

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
	routes.TrainRoutes(router)
	routes.ScheduleRoutes(router)

	router.Run("localhost:" + configs.EnvPORT())
}
