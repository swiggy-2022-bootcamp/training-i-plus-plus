package main

import (
	"io"
	"os"
	"sanitaria-microservices/patientModule/configs"
	"sanitaria-microservices/patientModule/routes"
	"sanitaria-microservices/patientModule/services"
	"github.com/gin-gonic/gin"
)

func main(){
	// Logging to a file.
    file,err := os.OpenFile("server.log", os.O_APPEND| os.O_CREATE | os.O_WRONLY, 0644)
    if err == nil{
		gin.DefaultWriter = io.MultiWriter(file)
	}
	
	router := gin.New()
	router.Use(services.UseLogger(services.DefaultLoggerFormatter), gin.Recovery())

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
				"data": "Server started successfully.",
		})
	})

	//connect database
	configs.ConnectDB()

	//routes
	routes.PatientRoutes(router)
	
	router.Run("localhost:8084") 
}