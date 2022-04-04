package main

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/taran1515/crud/configs"
	"github.com/taran1515/crud/docs"
	"github.com/taran1515/crud/routes"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("train_service.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
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

	docs.SwaggerInfo.Title = "Train Reservation System - Train Service"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	configs.ConnectDb()

	routes.TrainRoutes(router)

	router.Run("localhost:8001")

}
