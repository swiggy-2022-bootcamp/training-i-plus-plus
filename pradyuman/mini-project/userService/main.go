package main

import (
	
	"userService/configs"
	"userService/routes"
	// "mini-project/controllers"
	
	"github.com/gin-gonic/gin"	
)

func setupRouter() *gin.Engine {
	UserRouter := gin.Default()
	routes.UserRoute(UserRouter)
	return UserRouter
}


func main() {
	    //gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = configs.Logfile
		//setup database
		configs.ConnectDB()

        UserRouter := setupRouter()
		UserRouter.Run(":3000")
		
}