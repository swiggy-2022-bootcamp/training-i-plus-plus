package main

import (
	"github.com/gin-gonic/gin"
	"github.com/taran1515/crud/configs"
	"github.com/taran1515/crud/routes"
)

func main() {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"data": "Hello from Gin",
		})

	})

	configs.ConnectDb()

	routes.UserRoutes(router)

	router.Run("localhost:8000")

}
