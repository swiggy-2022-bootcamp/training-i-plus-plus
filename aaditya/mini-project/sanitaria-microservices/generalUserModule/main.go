package main

import (
	"fmt"
	"sanitaria-microservices/generalUserModule/configs"
	"sanitaria-microservices/generalUserModule/routes"
	"sanitaria-microservices/generalUserModule/services"

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
	routes.GeneralUserRoutes(router)
	
	router.Run("localhost:8083") 
}