package routes

import (
	"tejas/controllers"

	"github.com/gin-gonic/gin"
)

func HealthCheckRoutes(router *gin.Engine) {
	public := router.Group("/")
	public.GET("/health-check", controllers.HealthCheck())
	public.GET("/deep-health-check", controllers.DeepHealthCheck())
}
