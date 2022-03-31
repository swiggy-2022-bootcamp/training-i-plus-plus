package routes

import (
	"ticket_reservation_system/controller"
	"ticket_reservation_system/middleware"

	"github.com/gin-gonic/gin"
)

func TrainRoutes(router *gin.Engine) {
	trainRouter := router.Group("/train")
	trainRouter.Use(middleware.AuthenticateJWT())
	trainRouter.POST("/add", controller.AddTrain())
	trainRouter.GET("/get/:trainnumber", controller.GetTrainByTrainNumber())
	trainRouter.PUT("/update/:trainnumber", controller.UpdateTrainDetails())
	trainRouter.DELETE("/delete/:trainnumber", controller.DeleteTrain())

	router.GET("/trains", controller.GetAllTrains())
}
