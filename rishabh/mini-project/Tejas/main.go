package main

import (
	"tejas/configs"
	"tejas/routes"

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
	routes.TrainRoutes(router)

	router.Run("localhost:" + configs.EnvPORT())
}
