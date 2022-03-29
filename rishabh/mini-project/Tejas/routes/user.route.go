package routes

import (
	"tejas/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	public := router.Group("/api/user")
	public.POST("/register", controllers.RegisterUser())
	public.POST("/login", controllers.LoginUser())
}
