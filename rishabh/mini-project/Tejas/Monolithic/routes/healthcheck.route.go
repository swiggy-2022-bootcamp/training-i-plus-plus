package routes

import (
	"tejas/controllers"
	"tejas/middlewares"

	"github.com/gin-gonic/gin"
)

func HealthCheckRoutes(router *gin.Engine) {
	public := router.Group("/")
	public.Use(middlewares.LoggerMiddleware("HealthCheck"))

	public.GET("/health-check", controllers.HealthCheck())
	public.GET("/deep-health-check", controllers.DeepHealthCheck())
}
