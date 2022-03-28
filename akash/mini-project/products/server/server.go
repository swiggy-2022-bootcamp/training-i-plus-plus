package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"products.akash.com/api"
)

const (
	Port = "7778"
)

func Start() {

	// Configure
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello-world",
		})
	})

	router.GET("/product/:id", api.QueryOne)
	router.GET("/product/all", api.QueryAll)
	router.POST("/product/add", api.AddProduct)
	router.DELETE("/product/delete/:id", api.Delete)

	// Start listening
	port := Port
	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}

	fmt.Println("Server starting on port : ", port)
	router.Run(":" + port)
}
