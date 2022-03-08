package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

const (
	Port = "7777"
)

func Start() {

	// Configure
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello-world",
		})
	})

	// Start listening
	port := Port
	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}

	fmt.Println("Server starting on port : ", port)
	router.Run(":" + port)
}
