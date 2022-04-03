package main

import (
	"Order-Service/config"
	"Order-Service/controller"
	"Order-Service/middleware"
	"strconv"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.AllowAll())

	r.POST("/order", middleware.IfAuthorized(controller.PlaceOrder))
	r.GET("/:userId/order", middleware.IfAuthorized(controller.GetOrders))
	r.POST("/order/:orderId/payment", middleware.IfAuthorized(controller.OrderPayment))
	r.POST("/order/:orderId/deliver", middleware.IfAuthorized(controller.DeliverOrder))
	r.DELETE("/order/:orderId", middleware.IfAuthorized(controller.CancelOrder))

	portAddress := ":" + strconv.Itoa(config.ORDER_SERVICE_SERVER_PORT)
	r.Run(portAddress)
}
