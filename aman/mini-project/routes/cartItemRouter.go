package routes

import (
	"aman-swiggy-mini-project/controllers"
	"aman-swiggy-mini-project/middleware"

	"github.com/gin-gonic/gin"
)

func CartItemsRoutes(r *gin.Engine) {
	r.Use(middleware.Authentication())
	r.POST("/cart/add", controllers.AddCartItem())
	r.POST("/cart/:product_id", controllers.RemoveCartItem())
}
