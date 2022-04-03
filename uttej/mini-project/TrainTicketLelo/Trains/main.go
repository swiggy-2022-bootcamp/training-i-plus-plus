package main

import (
	"Trains/config"
	controllers "Trains/controller"
	"Trains/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/trains", middleware.IfAuthorized(controllers.CreateTrain))
	r.GET("/trains", middleware.IfAuthorized(controllers.GetTrains))
	r.GET("/trains/:trainId", middleware.IfAuthorized(controllers.GetTrainById))
	r.PUT("/trains/:trainId", middleware.IfAuthorized(controllers.UpdateTrainById))
	r.DELETE("/trains/:trainId", middleware.IfAuthorized(controllers.DeleteTrainbyId))
	r.POST("/trains/:trainId/:updateCount", middleware.IfAuthorized(controllers.UpdateTicketCount))

	portAddress := ":" + config.TrainServicePort
	r.Run(portAddress)

}
