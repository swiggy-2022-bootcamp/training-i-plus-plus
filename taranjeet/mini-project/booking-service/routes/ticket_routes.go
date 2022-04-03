package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/taran1515/crud/controllers"
)

func TicketRoutes(router *gin.Engine) {
	router.POST("/ticket/book", controllers.BookTicket())
	router.GET("/ticket/cancel/:ticketId", controllers.CancelBooking())
	router.PUT("/ticket/:ticketId", controllers.GetATicket())

}
