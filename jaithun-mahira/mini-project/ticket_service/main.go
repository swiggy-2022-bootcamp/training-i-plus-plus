package main

import (
	"ticket_service/configs"
	"ticket_service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()
  //run database
  configs.ConnectDB()

  router.GET("/", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "data": "Ticket Service",
    })
  })
  
  routes.TicketRoute(router)
  router.Run("localhost:6002") 
}