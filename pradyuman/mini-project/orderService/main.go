package main

import (
	
	"orderService/configs"
	"orderService/routes"
	"orderService/controllers"
	
	"github.com/gin-gonic/gin"	
)

func main() {
	    gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = configs.Logfile
		//setup database
		configs.ConnectDB()

		OrderRouter:= gin.Default()

		routes.OrderRoute(OrderRouter)
  
		go controllers.ProcessOrder()
		go controllers.UpdateOrderStatus()
		OrderRouter.Run(":3003")
}