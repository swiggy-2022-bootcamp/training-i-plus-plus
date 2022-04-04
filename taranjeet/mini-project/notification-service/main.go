package main

import (
	"github.com/gin-gonic/gin"
	"github.com/taran1515/crud/kafka"
)

func main() {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"data": "Hello from Gin",
		})

	})

	router.Run("localhost:8004")

	kafka.Consume()

}
