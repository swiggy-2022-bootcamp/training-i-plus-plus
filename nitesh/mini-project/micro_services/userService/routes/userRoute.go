package routes

import (
	"userService/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouter(gin *gin.Engine) {
	u := gin.Group("/user")
	{
		u.POST("/signup", controllers.Signup())
		u.POST("/login", controllers.Login())
	}
}
