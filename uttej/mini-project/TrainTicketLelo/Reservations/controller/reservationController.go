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

// BuyTicket godoc
// @Summary      Buy a Ticket
// @Description  Buy a Ticket by providing the details
// @Tags         Reservation
// @Accept       json
// @Produce      json
// @Param		Ticket	body	models.Reservation	true	"id, Departure Date, Purchase Date & Status Will Be Populated Automatically"
// @Success      200  {string}  responseBody
// @Failure      400  {number} 	http.StatusBadRequest
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /reservation [post]
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

// GetTickets godoc
// @Summary      Fetch All Tickets
// @Description  Get All Tickets for a user by providing the id
// @Tags         Reservation
// @Accept       json
// @Produce      json
// @Success      200  {array}  models.Reservation
// @Failure      400  {number} 	http.StatusBadRequest
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /:userId/reservations [get]
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

// TicketPayment godoc
// @Summary      Pay For a Ticket
// @Description  Pay For a ticket that you've reserved
// @Tags         Reservation
// @Accept       json
// @Produce      json
// @Param        TicketId 		body	string  true  "unique ticket id"
// @Success      200  {string}  successMessage
// @Failure      400  {number} 	http.StatusBadRequest
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /reservation/:ticketId/payment [post]
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

// CancelTicket godoc
// @Summary      Cancel a Ticket
// @Description  Cancel a Ticket that you've reserved
// @Tags         Reservation
// @Accept       json
// @Produce      json
// @Param        TicketId 		body	string  true  "unique ticket id"
// @Success      200  {string}  successMessage
// @Failure      400  {number} 	http.StatusBadRequest
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /reservation/:ticketId/cancel [post]
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
