package routes

import (
	"aman-swiggy-mini-project/controllers"
	"aman-swiggy-mini-project/middleware"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.Engine) {
	r.Use(middleware.Authentication())
	r.POST("/cart/purchase", controllers.CreateOrder())
	r.POST("/cancelOrder/:order_id", controllers.CancelOrder())
	r.POST("/orderPaid/:order_id", controllers.PaidOrder())
}
