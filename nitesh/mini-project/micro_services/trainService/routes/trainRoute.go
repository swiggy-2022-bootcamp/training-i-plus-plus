package routes

import (
	"trainService/controllers"

	"github.com/gin-gonic/gin"
)

func TrainRouter(gin *gin.Engine) {
	t := gin.Group("/train")
	{
		t.POST("/checkAvailability", controllers.CheckAvailability())
		t.GET("/searchRoute", controllers.SearchRoute())
		t.GET("/trainDetails", controllers.TrainDetails())
	}
}
