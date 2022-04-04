package main

import (
	
	"inventoryService/configs"
	"inventoryService/routes"
	
	"github.com/gin-gonic/gin"	
)

func main() {
	    //gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = configs.Logfile
		//setup database
		configs.ConnectDB()

		ProductRouter := gin.Default()
		
		routes.ProductViewRoute(ProductRouter)
		routes.ProductUpdateRoute(ProductRouter)
		
		ProductRouter.Run(":3001")
}