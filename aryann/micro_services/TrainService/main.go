package main

import (
	"TrainService/database"
	"TrainService/docs"
	"TrainService/logger"
	"TrainService/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

var logger1 = logger.NewLoggerService("main")

func TrainRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.TrainRoute(router)
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

// @securityDefinitions.basic  BasicAuth
func main() {

	database.ConnectDB()
	logger1.Log("Connected to MongoDB")

	trainroute := TrainRouter()
	// router := gin.Default()
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// routes.TrainRoute(router)
	docs.SwaggerInfo.Title = "Swagger Train Reservation System API"

	trainroute.Run("localhost:7000")
	// router.Run("localhost:7000")
}
