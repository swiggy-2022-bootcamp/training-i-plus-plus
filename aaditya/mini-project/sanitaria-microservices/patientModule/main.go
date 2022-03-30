package main

import (
	"io"
	"os"
	"sanitaria-microservices/patientModule/configs"
	"sanitaria-microservices/patientModule/routes"
	"sanitaria-microservices/patientModule/services"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "sanitaria-microservices/patientModule/docs"
)

// @title          Sanitaria - Patient Module
// @version        1.0
// @description    This microservice is for appointment module in the sanitaria application.
// @contact.name   Aaditya Khetan
// @contact.email  aadityakhetan123@gmail.com
// @license.name  Apache 2.0
// @host      localhost:8084
// @securityDefinitions.apikey  Bearer Token
// @in                          header
// @name                        Authorization

func main(){
	// Logging to a file.
    file,err := os.OpenFile("server.log", os.O_APPEND| os.O_CREATE | os.O_WRONLY, 0644)
    if err == nil{
		gin.DefaultWriter = io.MultiWriter(file)
	}
	
	router := gin.New()
	router.Use(services.UseLogger(services.DefaultLoggerFormatter), gin.Recovery())

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
				"data": "Server started successfully.",
		})
	})

	docs.SwaggerInfo.Title = "Sanitaria - Patient Module"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//connect database
	configs.ConnectDB()

	//routes
	routes.PatientRoutes(router)
	
	router.Run("localhost:8084") 
}