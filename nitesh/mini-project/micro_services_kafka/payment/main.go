package main

import (
	"paymentService/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.PaymentRoutes(router)

	router.Run(":8082")
}
