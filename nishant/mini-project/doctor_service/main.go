package main

import (
	"io"
	"log"
	"os"

	"github.com/swiggy-2022-bootcamp/training-i-plus-plus/nishant/mini-project/doctor_service/controller"
	"github.com/swiggy-2022-bootcamp/training-i-plus-plus/nishant/mini-project/doctor_service/db"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/swiggy-2022-bootcamp/training-i-plus-plus/nishant/mini-project/doctor_service/docs"
)

func init() {
	db.ConnectDB()
}

// @title Doctor Service
// @version 1.0
// @description This is doctor crud service.

// @host localhost:7451
// @BasePath /

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	f, err := os.OpenFile("doctor_service.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(io.MultiWriter(os.Stdout, f))
	log.Println("Logger setup!")

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PONG!",
		})
	})

	router.POST("/doctor", controller.Create)
	router.GET("/doctor", controller.Read)
	router.PATCH("/doctor/:_id", controller.Update)
	router.DELETE("/doctor/:_id", controller.Delete)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start listening
	router.Run(":7451")
}
