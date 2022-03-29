package main

import (
	"fmt"
	"sanitaria-microservices/appointmentModule/configs"
	"sanitaria-microservices/appointmentModule/routes"
	"sanitaria-microservices/appointmentModule/services"
	"github.com/gin-gonic/gin"
)

const consumerTopic = "Appointment"
func main(){
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
				"data": "Server started successfully.",
		})
	})

	//connect database
	configs.ConnectDB()
	consumer, err := services.CreateConsumer()
	if err != nil{
		fmt.Println("Error in creating kafka-consumer.")
	}else{
		go services.ConsumeAppointment(consumer,consumerTopic)
	}

	//routes
	routes.AppointmentRoutes(router)

	router.Run("localhost:8082") 
}