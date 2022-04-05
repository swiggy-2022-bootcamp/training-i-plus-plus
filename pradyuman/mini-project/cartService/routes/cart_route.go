package routes

import (
	"cartService/controllers"
	"cartService/middleware"
	"github.com/gin-gonic/gin"
)

func CartRoute(router *gin.Engine) {
	router.USE(middleware.IsUserAuthorized([]string{"BUYER"}))
	router.POST("/cart", controllers.AddItemToCart())
	router.GET("/cart/:userId", controllers.GetItemsfromCart())
	router.PUT("/cart/:cartId", controllers.EditItemFromCart())
	router.DELETE("/cart/:cartId", controllers.DeleteItemFromCart())
	router.GET("/placeOrder/:userId", controllers.PlaceOrderFromCart())
}
