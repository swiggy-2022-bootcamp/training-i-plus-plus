package main

import (
	"github.com/gin-gonic/gin"
	"sanitaria-microservices/generalUserModule/configs"
	"sanitaria-microservices/generalUserModule/routes"
)

func main(){
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
				"data": "Server started successfully.",
		})
	})

	//connect database
	configs.ConnectDB()

	//routes
	routes.GeneralUserRoutes(router)
	
	router.Run("localhost:8083") 
}