package main

import (
	"Order-Service/config"
	"Order-Service/controller"
	"Order-Service/middleware"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/order", middleware.IfAuthorized(controller.PlaceOrder))
	r.GET("/:userId/order", middleware.IfAuthorized(controller.GetOrders))
	r.POST("/order/:orderId/payment", middleware.IfAuthorized(controller.OrderPayment))
	r.POST("/order/:orderId/deliver", middleware.IfAuthorized(controller.DeliverOrder))

	portAddress := ":" + strconv.Itoa(config.ORDER_SERVICE_SERVER_PORT)
	r.Run(portAddress)
}
