package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"sample.akash.com/modules/user"
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

	router.POST("/register", user.Register)
	router.GET("/login", user.Login)
	router.GET("/user", user.QueryOne)
	router.GET("/user/all", user.QueryAll)
	router.PUT("/user/update", user.Update)
	router.DELETE("/user/delete", user.Delete)

	// Start listening
	port := Port
	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}

	fmt.Println("Server starting on port : ", port)
	router.Run(":" + port)
}
