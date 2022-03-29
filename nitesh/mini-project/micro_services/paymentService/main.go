package main

import (
	"paymentService/routes"

	"paymentService/docs"

	"github.com/gin-gonic/gin" // swagger embed files
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title           Swagger Train-Ticket Booking System API
// @version         1.0
// @description     Swagger Train-Ticket Booking System API.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  swiggyb2013@datascience.manipal.edu

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8082

// @securityDefinitions.basic  BasicAuth
func main() {
	router := gin.Default()
	docs.SwaggerInfo.Title = "Swagger Train-Ticket Booking System API"

	routes.PaymentRoutes(router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8082")
}
