package main

import (
	"github.com/gin-gonic/gin"

	"usecase/user_crud_dynamo/controller"
	"usecase/user_crud_dynamo/db"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PONG!",
		})
	})

	_, svc := db.Connect()

	cont := controller.Controller{svc}

	router.POST("/user", cont.CreateUser)
	// router.GET("/doctor", handler.Read)
	// router.PATCH("/doctor/:_id", handler.Update)
	// router.DELETE("/doctor/:_id", handler.Delete)

	// Start listening
	router.Run(":7456")
}
