package routes

import (
	"aman-swiggy-mini-project/controllers"
	"aman-swiggy-mini-project/middleware"

	"github.com/gin-gonic/gin"
)

func CartRoutes(r *gin.Engine) {
	r.Use(middleware.Authentication())
	r.GET("/cart/:cart_id", controllers.GetProduct())
	r.POST("/cart", controllers.CreateProduct())
}
