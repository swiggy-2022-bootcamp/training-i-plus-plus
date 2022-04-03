package routes

import (
	"aman-swiggy-mini-project/controllers"
	"aman-swiggy-mini-project/middleware"

	"github.com/gin-gonic/gin"
)

func RequestRoutes(r *gin.Engine) {
	r.Use(middleware.Authentication())
	r.GET("/request", controllers.GetRequest())
	r.POST("/request", controllers.PostRequest())
}
