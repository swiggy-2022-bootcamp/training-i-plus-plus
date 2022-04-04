package main

import (
	"auth/configs"
	"auth/routes"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	AuthRouter := gin.Default()
	routes.AuthRoute(AuthRouter)
	return AuthRouter
}

func main() {
	    gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = configs.Logfile
		//setup database
		configs.ConnectDB()

        AuthRouter := setupRouter()
		AuthRouter.Run(":3005")
		
}