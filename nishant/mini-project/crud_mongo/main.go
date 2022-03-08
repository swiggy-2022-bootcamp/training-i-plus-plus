package main

import (
	"usecase/crud_mongo/db"
	"usecase/crud_mongo/handler"

	"github.com/gin-gonic/gin"
)

func init() {
	db.ConnectDB()
}

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PONG!",
		})
	})

	router.POST("/doctor", handler.Create)
	router.GET("/doctor", handler.Read)
	router.PATCH("/doctor/:_id", handler.Update)
	router.DELETE("/doctor/:_id", handler.Delete)

	// Start listening
	router.Run(":7456")
}
