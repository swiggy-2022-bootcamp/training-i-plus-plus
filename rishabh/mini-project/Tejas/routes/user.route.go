package routes

import (
	"tejas/controllers"
	"tejas/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	public := router.Group("/api/user")
	public.Use(middlewares.LoggerMiddleware("User"))

	public.POST("/register", controllers.RegisterUser())
	public.POST("/login", controllers.LoginUser())
}
