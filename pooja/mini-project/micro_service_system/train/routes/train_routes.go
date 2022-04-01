package routes

import (
	"train/middleware"
	"train/service"

	"github.com/gin-gonic/gin"
)

func TrainRoutes(router *gin.Engine) {
	trainRouter := router.Group("/train")
	trainRouter.Use(middleware.AuthenticateJWT())
	trainRouter.POST("/add", service.AddTrain())
	trainRouter.GET("/get/:trainnumber", service.GetTrainByTrainNumber())
	trainRouter.PUT("/update/:trainnumber", service.UpdateTrainDetails())
	trainRouter.DELETE("/delete/:trainnumber", service.DeleteTrain())

	router.GET("/trains", service.GetAllTrains())
	router.GET("/checkavailability", service.CheckSeatAvailablity())
	router.GET("/search_trains", service.SearchTrains())

}
