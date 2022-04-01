package routes

import (
	"reservation/middleware"
	"reservation/service"

	"github.com/gin-gonic/gin"
)

func ReservationRoutes(router *gin.Engine) {
	reservationRouter := router.Group("/reservation")
	reservationRouter.Use(middleware.AuthenticateJWT())
	reservationRouter.POST("/reserve_ticket", service.BookTickets())
	reservationRouter.PUT("/cancel_reservation", service.CancelBooking())
	reservationRouter.GET("/bookings", service.ViewBookings())

}
