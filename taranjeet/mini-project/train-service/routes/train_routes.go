package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/taran1515/crud/controllers"
)

func TrainRoutes(router *gin.Engine) {
	router.POST("/train/add", controllers.CreateTrain())
	router.GET("/train/:trainId", controllers.GetATrain())
	router.PUT("/train/:trainId", controllers.EditATrain())
	router.DELETE("/train/:trainId", controllers.DeleteATrain())
	router.GET("/train", controllers.GetAllTrains())
	router.POST("/train/search", controllers.SearchTrain())

}
