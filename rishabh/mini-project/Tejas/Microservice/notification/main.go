package main

import (
	"notificationService/configs"
	"notificationService/kafka"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server is up and running",
		})
	})

	go kafka.PaymentDetails()
	go kafka.TicketDetails()
	router.Run("localhost:" + configs.EnvPORT())
}
