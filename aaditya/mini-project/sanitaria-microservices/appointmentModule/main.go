package main

import (
	"github.com/gin-gonic/gin"
	"sanitaria-microservices/appointmentModule/configs"
	"sanitaria-microservices/appointmentModule/routes"
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
	routes.AppointmentRoutes(router)

	router.Run("localhost:8082") 
}