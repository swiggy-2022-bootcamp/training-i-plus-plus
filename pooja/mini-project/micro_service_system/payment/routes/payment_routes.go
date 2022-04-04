package routes

import (
	"payment/middleware"
	"payment/service"

	"github.com/gin-gonic/gin"
)

func PaymentRoutes(router *gin.Engine) {
	paymentRouter := router.Group("/payment")
	paymentRouter.Use(middleware.AuthenticateJWT())
	paymentRouter.POST("/pay", service.Payment())
}
