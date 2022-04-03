package controller

import (
	repository "Order-Service/Repository"
	"Order-Service/errors"
	mockdata "Order-Service/model"
	service "Order-Service/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var orderService service.IOrderService

func init() {
	orderService = service.InitOrderService(&repository.MongoDAO{}, &repository.HttpRepo{})
}

func PlaceOrder(c *gin.Context) {
	//access: Anyone

	var orderPlaced mockdata.Order
	json.NewDecoder(c.Request.Body).Decode(&orderPlaced)
	result, error := orderService.PlaceOrder(orderPlaced)

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

	orders, error := orderService.GetOrders(userId)

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
	successMessage, error := orderService.OrderPayment(orderId)

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
	successMessage, error := orderService.DeliverOrder(orderId)

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

func CancelOrder(c *gin.Context) {
	//access: a user can cancel his/her own undelivered orders
	//acessorUserRole, _ := strconv.Atoi(c.Param("acessorUserRole"))
	acessorUserId := c.Param("acessorUserId")
	orderId := c.Param("orderId")

	msg, err := orderService.CancelOrder(orderId, acessorUserId)
	if err != nil {
		orderError, ok := err.(*errors.OrderError)
		if ok {
			c.JSON(orderError.Status, orderError.ErrorMessage)
			return
		} else {
			fmt.Println("orderError casting error in CancelOrder")
			return
		}
	}

	c.JSON(http.StatusOK, *msg)
}

//confirmed -> payment done -> delivered
