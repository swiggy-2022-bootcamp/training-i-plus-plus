package main

import (
	"Reservations/config"
	"Reservations/controller"
	"Reservations/docs"
	"Reservations/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Swagger TrainTicketLelo Reservations Service
// @version         1.0
// @description     Swagger TrainTicketLelo Reservations Service
// @termsOfService  http://swagger.io/terms/

// @contact.name   Uttej Immadi
// @contact.url    http://www.swagger.io/support
// @contact.email  immadiuttej@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8002

// @securityDefinitions.apiKey ApiKeyAuth
// @type apiKey
// @in header
// @name Authorization
func main() {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/reservation", middleware.IfAuthorized(controller.BuyTicket))
	r.GET("/:userId/reservations", middleware.IfAuthorized(controller.GetTickets))
	r.POST("/reservation/:ticketId/payment", middleware.IfAuthorized(controller.TicketPayment))
	r.POST("/reservation/:ticketId/cancel", middleware.IfAuthorized(controller.CancelTicket))

	docs.SwaggerInfo.Title = "Swagger TrainTicketLelo Reservations Service"

	portAddress := ":" + config.ReservationServicePort
	r.Run(portAddress)
}
