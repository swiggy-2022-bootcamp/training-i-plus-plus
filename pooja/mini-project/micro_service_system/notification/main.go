package main

import (
	"notification/kafka"

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
	router.Run("localhost:6005")
}
