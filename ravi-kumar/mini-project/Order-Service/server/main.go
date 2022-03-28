package main

import (
	"Order-Service/config"
	"Order-Service/controller"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/order", controller.PlaceOrder)
	r.GET("/:userId/order", controller.GetOrders)
	r.POST("/order/:orderId/payment", controller.OrderPayment)
	r.POST("/order/:orderId/deliver", controller.DeliverOrder)

	portAddress := ":" + strconv.Itoa(config.ORDER_SERVICE_SERVER_PORT)
	r.Run(portAddress)
}
