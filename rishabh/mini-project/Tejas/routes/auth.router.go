package routes

import (
	"tejas/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(gin *gin.Engine) {

	auth := gin.Group("/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/signup", controllers.Signup)
	}
}
