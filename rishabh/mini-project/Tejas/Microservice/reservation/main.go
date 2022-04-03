package main

import (
	"reservationService/configs"
	"reservationService/routes"
	"reservationService/services"

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
	routes.ScheduleRoutes(router)

	router.Run("localhost:" + configs.EnvPORT())
}
