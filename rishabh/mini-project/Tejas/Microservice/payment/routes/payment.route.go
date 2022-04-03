package routes

import (
	"paymentService/controllers"
	"paymentService/middlewares"

	"github.com/gin-gonic/gin"
)

func PaymentRoutes(router *gin.Engine) {
	private := router.Group("/api/payment")
	private.Use(middlewares.LoggerMiddleware("Payment"))
	private.Use(middlewares.AuthenticateJWT())
	private.POST("/pay", controllers.Payment())
}
