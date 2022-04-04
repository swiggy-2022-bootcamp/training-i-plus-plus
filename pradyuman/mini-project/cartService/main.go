package main

import (
	
	"cartService/configs"
	"cartService/routes"
	// "mini-project/controllers"
	
	"github.com/gin-gonic/gin"	
)

func main() {
	    gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = configs.Logfile
		//setup database
		configs.ConnectDB()

		CartRouter := gin.Default()
	
		routes.CartRoute(CartRouter)

		CartRouter.Run(":3002")
		
}