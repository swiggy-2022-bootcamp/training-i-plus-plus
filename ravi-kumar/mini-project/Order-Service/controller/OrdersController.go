package controller

import (
	"Order-Service/errors"
	mockdata "Order-Service/model"
	service "Order-Service/service"
	"fmt"
	"net/http"
	"strconv"

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
	//access: Admin and Customer
	acessorUserRole, _ := strconv.Atoi(c.Param("acessorUserRole"))
	acessorUserId := c.Param("acessorUserId")
	userId := c.Param("userId")

	if !((mockdata.Role(acessorUserRole) == mockdata.Customer && acessorUserId == userId) || mockdata.Role(acessorUserRole) == mockdata.Admin) {
		c.JSON(http.StatusUnauthorized, errors.AccessDenied())
		return
	}

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
	//access: Admin and Customer
	acessorUserRole, _ := strconv.Atoi(c.Param("acessorUserRole"))
	if !(mockdata.Role(acessorUserRole) == mockdata.Customer || mockdata.Role(acessorUserRole) == mockdata.Admin) {
		c.JSON(http.StatusUnauthorized, errors.AccessDenied())
		return
	}

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
	//access: Admin
	acessorUserRole, _ := strconv.Atoi(c.Param("acessorUserRole"))
	if mockdata.Role(acessorUserRole) != mockdata.Admin {
		c.JSON(http.StatusUnauthorized, errors.AccessDenied())
		return
	}

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
