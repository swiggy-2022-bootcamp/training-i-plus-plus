package main

import (
	"io"
	
	
	"log"
	routes "bhatiachahat/order-service/routes"
	utils "bhatiachahat/order-service/utils"
 //  kafkaservice "bhatiachahat/order-service/kafkaservice"
    "os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "bhatiachahat/order-service/docs"




)
// @title          Online Shopping Application - Orders Module
// @version        1.0
// @description    This microservice is for orders module in the online shopping application.
// @contact.name   Chahat Bhatia
// @contact.email  chahatbhatia2014@gmail.com
// @license.name  Apache 2.0
// @host          localhost:8082
// @securityDefinitions.apikey  Bearer Token
// @in                          header
// @name                        Authorization
func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	if port==""{
		port="8082"}

		// Logging to a file.
		file,err := os.OpenFile("server.log", os.O_APPEND| os.O_CREATE | os.O_WRONLY, 0644)
		if err == nil{
			gin.DefaultWriter = io.MultiWriter(file)
		}
		
		router := gin.New()
		router.Use(gin.Logger())
		router.Use(utils.UseLogger(utils.DefaultLoggerFormatter), gin.Recovery())
		
		router.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
					"data": "Server started successfully.",
			})
		})
	
		docs.SwaggerInfo.Title = "Online Shopping App - Orders Module"
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
		
		//routes
		//kafkaservice.StartKafkaConsumer()

	    routes.OrderRoutes(router)
    	router.Run(":" + port)

}