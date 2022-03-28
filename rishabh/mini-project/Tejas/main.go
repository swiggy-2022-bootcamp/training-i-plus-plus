package main

import (
	"authApp/configs"
	"authApp/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server is up and running",
		})
	})

	// init db
	configs.ConnectDB()

	// init routes
	routes.UserRoutes(router)

	router.Run("localhost:" + configs.EnvPORT())
}
