package controller

import (
	"Order-Service/errors"
	service "Order-Service/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PlaceOrder(c *gin.Context) {
	result, error := service.PlaceOrder(&c.Request.Body)

	if error != nil {
		orderError, ok := error.(*errors.OrderError)
		if ok {
			c.JSON(orderError.Status, orderError.ErrorMessage)
			return
		} else {
			fmt.Println("orderError casting error in PlaceOrder")
			return
		}
	}
	c.JSON(http.StatusOK, result)
}

func GetOrders(c *gin.Context) {
	userId := c.Param("userId")
	orders, error := service.GetOrders(userId)

	if error != nil {
		orderError, ok := error.(*errors.OrderError)
		if ok {
			c.JSON(orderError.Status, orderError.ErrorMessage)
			return
		} else {
			fmt.Println("orderError casting error in GetOrders")
			return
		}
	}
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
			fmt.Println("orderError casting error in DeliverOrder")
			return
		}
	}
	c.JSON(http.StatusOK, *successMessage)
}

//confirmed -> payment done -> delivered
