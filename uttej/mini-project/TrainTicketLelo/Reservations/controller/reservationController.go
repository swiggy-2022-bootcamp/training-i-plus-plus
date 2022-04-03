package controller

import (
	"Reservations/errors"
	models "Reservations/model"
	service "Reservations/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func BuyTicket(c *gin.Context) {
	result, error := service.BuyTicket(&c.Request.Body)

	if error != nil {
		purchaseError, ok := error.(*errors.PurchaseError)
		if ok {
			c.JSON(purchaseError.Status, purchaseError.ErrorMessage)
			return
		} else {
			fmt.Println("Error Buying The Ticket")
			return
		}
	}
	c.JSON(http.StatusOK, result)
}

func GetTickets(c *gin.Context) {
	acessorUserRole, _ := strconv.Atoi(c.Param("acessorUserRole"))
	acessorUserId := c.Param("acessorUserId")
	userId := c.Param("userId")

	if !((models.Role(acessorUserRole) == models.Traveller && acessorUserId == userId) || models.Role(acessorUserRole) == models.Admin) {
		c.JSON(http.StatusUnauthorized, errors.AccessDenied())
		return
	}

	purchases, error := service.GetTickets(userId)

	if error != nil {
		purchaseError, ok := error.(*errors.PurchaseError)
		if ok {
			c.JSON(purchaseError.Status, purchaseError.ErrorMessage)
			return
		} else {
			fmt.Println("Couldn't Fetch Tickets")
			return
		}
	}
	c.JSON(http.StatusOK, purchases)
}

func TicketPayment(c *gin.Context) {
	acessorUserRole, _ := strconv.Atoi(c.Param("acessorUserRole"))
	if !(models.Role(acessorUserRole) == models.Traveller || models.Role(acessorUserRole) == models.Admin) {
		c.JSON(http.StatusUnauthorized, errors.AccessDenied())
		return
	}

	ticketId := c.Param("ticketId")
	successMessage, error := service.TicketPayment(ticketId)

	if error != nil {
		purchaseError, ok := error.(*errors.PurchaseError)
		if ok {
			c.JSON(purchaseError.Status, purchaseError.ErrorMessage)
			return
		} else {
			fmt.Println("Couldn't Finish Payment Process")
			return
		}
	}
	c.JSON(http.StatusOK, *successMessage)
}

func CancelTicket(c *gin.Context) {
	acessorUserRole, _ := strconv.Atoi(c.Param("acessorUserRole"))
	if !(models.Role(acessorUserRole) == models.Traveller || models.Role(acessorUserRole) == models.Admin) {
		c.JSON(http.StatusUnauthorized, errors.AccessDenied())
		return
	}

	ticketId := c.Param("ticketId")
	successMessage, error := service.CancelTicket(ticketId)

	if error != nil {
		purchaseError, ok := error.(*errors.PurchaseError)
		if ok {
			c.JSON(purchaseError.Status, purchaseError.ErrorMessage)
			return
		} else {
			fmt.Println("Couldn't Finish Cancelling Process")
			return
		}
	}
	c.JSON(http.StatusOK, *successMessage)
}
