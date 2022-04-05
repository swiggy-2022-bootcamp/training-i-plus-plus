package routes

import (
	"orderService/controllers"
	"orderService/middleware"
	"github.com/gin-gonic/gin"
)

func OrderRoute(router *gin.Engine) {
	router.Use(middleware.IsUserAuthorized([]string{"BUYER","SELLER"}))
	router.GET("/orders/user/:userId", controllers.GetAllUserOrders())
	router.GET("/orders/seller/:sellerId", controllers.GetAllSellerOrders())
	//router.PUT("/order/:orderId", controllers.EditOrder())
}
