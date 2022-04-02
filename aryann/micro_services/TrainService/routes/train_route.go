package routes

import (
	"TrainService/controllers"

	"github.com/gin-gonic/gin"
)

func TrainRoute(router *gin.Engine) {
	router.Use(controllers.IsAuthorized("admin"))
	router.POST("/train", controllers.CreateTrain())
	router.GET("/train/:trainid", controllers.GetTrain())
	router.PUT("/train/:trainid", controllers.EditTrain())
	router.DELETE("/train/:trainid", controllers.DeleteTrain())
}
