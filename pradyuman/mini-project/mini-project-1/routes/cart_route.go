package routes

import (
	"mini-project/controllers"

	"github.com/gin-gonic/gin"
)

func CartRoute(router *gin.Engine) {
	router.POST("/cart", controllers.AddItemToCart())
	router.GET("/cart/:userId", controllers.GetItemsfromCart())
	router.PUT("/cart/:cartId", controllers.EditItemFromCart())
	router.DELETE("/cart/:cartId", controllers.DeleteItemFromCart())
	router.GET("/placeOrder/:userId", controllers.PlaceOrderFromCart())
}
