package routes

import (
	"ticket_service/controllers"
	"ticket_service/middlewares"

	"github.com/gin-gonic/gin"
)

func TicketRoute(router *gin.Engine)  {

	router.Use(middlewares.Authenticate("ADMIN", false))
	router.GET("/tickets", controllers.GetAllTickets())

	userRouter := router.Group("/tickets")
	userRouter.Use(middlewares.Authenticate("USER", true))
	userRouter.POST("/", controllers.BookTicket())
	userRouter.PUT("/:ticketId/cancelTicket", controllers.CancelTicket())

	adminRouter := router.Group("/tickets")
	adminRouter.Use(middlewares.Authenticate("ADMIN", true))
	adminRouter.GET("/:ticketId", controllers.GetTicketById())

}