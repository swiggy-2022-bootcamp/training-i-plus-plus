package main

import (
	"tejas/configs"
	"tejas/docs"
	"tejas/routes"
	"tejas/services"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var logger = services.NewLoggerService("main")

// @title           Swagger Tejas API
// @version         1.0
// @description     Swagger Train Reservation System API.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Rishabh Mishra
// @contact.url    htpps://rishabhmishra.me
// @contact.email  swiggyb2026@datascience.manipal.edu

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3000

// @securityDefinitions.basic  BasicAuth
func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server is up and running",
		})
	})

	// init db
	configs.ConnectDB()
	logger.Log("Connected to DB")

	// init routes
	routes.UserRoutes(router)
	routes.TrainRoutes(router)
	routes.ScheduleRoutes(router)
	routes.HealthCheckRoutes(router)

	// swagger
	docs.SwaggerInfo.Title = "Tejas API"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run("localhost:" + configs.EnvPORT())
}
