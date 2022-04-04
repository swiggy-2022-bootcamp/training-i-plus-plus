package routes

import (
	"train_service/controllers"
	"train_service/middlewares"

	"github.com/gin-gonic/gin"
)

func TrainRoute(router *gin.Engine)  {
	router.Use(middlewares.Authenticate("ADMIN", false))
	router.GET("/trains", controllers.GetAllTrains())
	router.GET("/trains/:trainId", controllers.GetTrainById())
    
	adminRouter := router.Group("/trains")
	adminRouter.Use(middlewares.Authenticate("ADMIN", true))
	adminRouter.POST("/", controllers.AddTrain())
	adminRouter.PUT("/:trainId", controllers.UpdateTrain())
	adminRouter.DELETE("/:trainId", controllers.DeleteTrain())

}