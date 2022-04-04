package main

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/swiggy-2022-bootcamp/training-i-plus-plus/nishant/mini-project/user_service/controller"
	"github.com/swiggy-2022-bootcamp/training-i-plus-plus/nishant/mini-project/user_service/db"
	"github.com/swiggy-2022-bootcamp/training-i-plus-plus/nishant/mini-project/user_service/producer"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/swiggy-2022-bootcamp/training-i-plus-plus/nishant/mini-project/user_service/docs"
)

// @title User Service
// @version 1.0
// @description This is user crud service.

// @host localhost:7450
// @BasePath /

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	f, err := os.OpenFile("user_service.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
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

	_, svc := db.Connect()
	p := producer.NewProducer("test_topic")
	cont := controller.Controller{svc, p}

	router.POST("/user", cont.CreateUser, cont.CreateToken)

	router.GET("/user/:_id", cont.ReadUser)
	router.PATCH("/user/:_id", cont.VerifyToken, cont.UpdateUser)
	router.DELETE("/user/:_id", cont.VerifyToken, cont.DeleteUser)
	router.GET("/user", cont.VerifyToken, cont.ListUser)

	router.POST("/login", cont.Login, cont.CreateToken)

	router.GET("/appointments", cont.VerifyToken, cont.GetMyAppointments)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start listening
	router.Run(":7450")
}
