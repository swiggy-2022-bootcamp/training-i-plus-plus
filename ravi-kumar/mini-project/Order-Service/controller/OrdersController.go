package controller

import (
	"Order-Service/errors"
	service "Order-Service/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PlaceOrder(c *gin.Context) {
	result := service.PlaceOrder(&c.Request.Body)
	c.JSON(http.StatusOK, result)
}

func GetOrders(c *gin.Context) {
	userId := c.Param("userId")
	orders := service.GetOrders(userId)
	c.JSON(http.StatusOK, orders)
}

func OrderPayment(c *gin.Context) {
	orderId := c.Param("orderId")
	successMessage, error := service.OrderPayment(orderId)

	if error != nil {
		orderError, ok := error.(*errors.OrderError)
		if ok {
			c.JSON(orderError.Status, orderError.ErrorMessage)
			return
		} else {
			fmt.Println("orderError casting error in OrderPayment")
			return
		}
	}
	c.JSON(http.StatusOK, *successMessage)
}

func DeliverOrder(c *gin.Context) {
	orderId := c.Param("orderId")
	successMessage, error := service.DeliverOrder(orderId)

	if error != nil {
		orderError, ok := error.(*errors.OrderError)
		if ok {
			c.JSON(orderError.Status, orderError.ErrorMessage)
			return
		} else {
			fmt.Println("orderError casting error in OrderPayment")
			return
		}
	}
	c.JSON(http.StatusOK, *successMessage)
}

//confirmed -> payment done -> delivered
