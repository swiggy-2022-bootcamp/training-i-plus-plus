package main

import (
	"github.com/gin-gonic/gin"
	"sanitaria/configs"
	"sanitaria/routes"
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
	routes.DoctorRoutes(router)
	routes.PatientRoutes(router)
	

	router.Run("localhost:8082") 
}