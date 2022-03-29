package main

import (
	"fmt"
	"sanitaria-microservices/doctorModule/configs"
	"sanitaria-microservices/doctorModule/routes"
	"sanitaria-microservices/doctorModule/services"

	"github.com/gin-gonic/gin"
)
const(
	consumerTopic = "Booked-appointment"
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

	consumer, err := services.CreateConsumer()
	if err != nil{
		fmt.Println("Error in creating kafka-consumer.")
	}else{
		go services.ConsumeBookedAppointment(consumer,consumerTopic)
	}

	//routes
	routes.DoctorRoutes(router)

	router.Run("localhost:8081") 
}