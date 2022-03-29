package main

import (
	"github.com/gin-gonic/gin"
	"sanitaria-microservices/patientModule/configs"
	"sanitaria-microservices/patientModule/routes"
)

func main(){
	router := gin.Default()

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