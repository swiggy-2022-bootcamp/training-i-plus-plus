package main

import (
	"reservation/database"
	"reservation/docs"
	"reservation/routes"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:6003
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func main() {
	router := gin.Default()

	docs.SwaggerInfo.Title = "Train ticket reservation system - Reservation module"

	database.DatabaseConn()
	log.Info("Reservation module is connected to db")

	routes.ReservationRoutes(router)
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":6003")
}
