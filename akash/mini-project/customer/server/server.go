package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"sample.akash.com/api"
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

	router.POST("/register", api.Register)
	router.GET("/login", api.Login)
	router.GET("/user", api.QueryOne)
	router.GET("/user/all", api.QueryAll)
	router.PUT("/user/update", api.Update)
	router.DELETE("/user/delete", api.Delete)

	// Start listening
	port := Port
	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}

	fmt.Println("Server starting on port : ", port)
	router.Run(":" + port)
}
