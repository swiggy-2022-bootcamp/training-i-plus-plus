package routes

import (
	"TicketService/controllers"

	"github.com/gin-gonic/gin"
)

func TicketRoute(router *gin.Engine) {
	router.Use(controllers.IsAuthorized("user"))
	router.POST("/ticket", controllers.CreateTicket())
	router.GET("/ticket/:ticketid", controllers.GetTicket())
	router.DELETE("/ticket/:ticketid", controllers.DeleteTicket())
}
