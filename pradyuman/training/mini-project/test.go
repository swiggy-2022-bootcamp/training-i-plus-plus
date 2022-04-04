package main

import (
	
	"mini-project/configs"
	"mini-project/routes"
	
	"github.com/gin-gonic/gin"	
)

func main() {
	    //gin.SetMode(gin.ReleaseMode)
		
        router := gin.Default()
  
		//setup database
		configs.ConnectDB()

		//routes
		routes.UserRoute(router)

        router.GET("/", func(c *gin.Context) {
                c.JSON(200, gin.H{
                        "data": "Hello from Gin-gonic and mongoDB",
                })
        })
		
		router.Run(":"+configs.EnvPort())
}