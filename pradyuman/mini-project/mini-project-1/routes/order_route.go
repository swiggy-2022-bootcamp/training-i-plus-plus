package routes

import (
	"mini-project/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoute(router *gin.Engine) {
	//router.POST("/cart", controllers.AddItemToCart())
	router.GET("/order/:userId", controllers.GetAllOrders())
	// router.PUT("/cart/:cartId", controllers.EditItemFromCart())
	// router.DELETE("/cart/:cartId", controllers.DeleteItemFromCart())
	// router.DELETE("/placeOrder/:userId", controllers.PlaceOrderFromCart())
}
