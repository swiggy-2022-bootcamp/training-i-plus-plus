package main

import (
	"TicketService/database"
	"TicketService/docs"
	"TicketService/logger"
	"TicketService/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

var logger1 = logger.NewLoggerService("main")

func TicketRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.TicketRoute(router)
	return router
}

// @title           Swagger Train Reservation System API
// @version         1.0
// @description     Swagger Train Reservation System API
// @termsOfService  http://swagger.io/terms/

// @contact.name   Aryann Dhir
// @contact.url    http://www.swagger.io/support
// @contact.email  swiggyb3053@datascience.manipal.edu

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:7000

// @securityDefinitions.apiKey ApiKeyAuth
// @type apiKey
// @in header
// @name Authorization
func main() {

	database.ConnectDB()
	logger1.Log("Connected to MongoDB")

	docs.SwaggerInfo.Title = "Swagger Train Reservation System API"

	userroute := TicketRouter()
	userroute.Run("localhost:7000")
}
