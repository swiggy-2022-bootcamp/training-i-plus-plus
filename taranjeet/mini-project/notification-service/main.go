package main

import (
	"github.com/gin-gonic/gin"
	"github.com/taran1515/crud/kafka"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("notification_service.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(io.MultiWriter(os.Stdout, f))
	log.Println("Logger setup!")

	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"data": "Hello from Gin",
		})

	})

	router.Run("localhost:8004")

	kafka.Consume()

}
