package routes

import (
	"golang-trainticketlelo/controllers",
	"gitub.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("/users/register", controllers.Register())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
	incomingRoutes.GET("/users", controllers.GetUsers())
}