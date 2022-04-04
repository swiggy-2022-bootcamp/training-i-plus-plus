package main

import (
	"user_service/configs"
	"user_service/logger"
	"user_service/routes"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
  log := logger.InitializeLogger()

  zap.ReplaceGlobals(log)
  defer log.Sync()

  log.Info("User Service Application Started")
  router := gin.Default()
  //run database
  configs.ConnectDB()

  router.GET("/", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "data": "User Service",
    })
  })
  
  routes.UserRoute(router)
  router.Run("localhost:6001") 
}