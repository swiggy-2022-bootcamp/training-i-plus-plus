package main

import (
	"Reservations/config"
	"Reservations/controller"
	"Reservations/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/reservation", middleware.IfAuthorized(controller.BuyTicket))
	r.GET("/:userId/reservations", middleware.IfAuthorized(controller.GetTickets))
	r.POST("/reservation/:ticketId/payment", middleware.IfAuthorized(controller.TicketPayment))

	portAddress := ":" + config.ReservationServicePort
	r.Run(portAddress)
}
