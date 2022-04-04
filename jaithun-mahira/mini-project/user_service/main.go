package main

import (
	"user_service/configs"
	_ "user_service/docs"
	"user_service/logger"
	"user_service/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

// @title           Swagger User Service API
// @version         1.0
// @description     Contains User APIs for Train Reservation System
// @termsOfService  http://swagger.io/terms/

// @contact.name   Jaithun Mahira
// @contact.email  swiggyb1035@datascience.manipal.edu

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:6001

// @securityDefinitions.apiKey ApiKeyAuth
// @type apiKey
// @in header
// @name Authorization
func main() {
  log := logger.InitializeLogger()

  zap.ReplaceGlobals(log)
  defer log.Sync()

  log.Info("User Service Application Started")
  //run database
  configs.ConnectDB()

  router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

  router.GET("/", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "data": "User Service",
    })
  })
  
  routes.UserRoute(router)
  router.Run("localhost:6001") 
}