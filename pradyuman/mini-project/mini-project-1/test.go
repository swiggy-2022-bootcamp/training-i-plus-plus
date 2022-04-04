package main

import (
	
	"mini-project/configs"
	"mini-project/routes"
	// "mini-project/controllers"
	
	"github.com/gin-gonic/gin"	
)

func main() {
	    gin.SetMode(gin.ReleaseMode)
		
		//setup database
		configs.ConnectDB()


        UserRouter := gin.Default()
		ProductRouter := gin.Default()
		CartRouter := gin.Default()
		OrderRouter:= gin.Default()

		//Services
		routes.UserRoute(UserRouter)
		routes.ProductRoute(ProductRouter)
		routes.CartRoute(CartRouter)
		routes.OrderRoute(OrderRouter)
  
        
		//go controllers.ProcessOrder()
		go UserRouter.Run(":3000")
		go ProductRouter.Run(":3001")
		go OrderRouter.Run(":3004")
		CartRouter.Run(":3002")
		
}