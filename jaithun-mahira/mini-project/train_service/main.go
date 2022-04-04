package main

import (
	"train_service/configs"
	"train_service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()
  //run database
  configs.ConnectDB()

  router.GET("/", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "data": "Train Service",
    })
  })
  
  routes.TrainRoute(router)
  router.Run("localhost:6000") 
}