package routes

import (
	"userService/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouter(gin *gin.Engine) {
	t := gin.Group("/user")
	{
		t.POST("/book", controllers.BookTicket())
	}
}
