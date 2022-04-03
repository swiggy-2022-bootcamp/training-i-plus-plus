package main

import (
	"io"
	
	
	"log"
	routes "bhatiachahat/product-service/routes"
	utils "bhatiachahat/product-service/utils"
   kafkaservice "bhatiachahat/product-service/kafkaservice"
    "os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "bhatiachahat/product-service/docs"




)
// @title          Online Shopping Application - Products Module
// @version        1.0
// @description    This microservice is for doctor module in the online shopping application.
// @contact.name   Chahat Bhatia
// @contact.email  chahatbhatia2014@gmail.com
// @license.name  Apache 2.0
// @host          localhost:8080
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
		port="8080"}

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
	
		docs.SwaggerInfo.Title = "Online Shopping App - Products Module"
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
		
		//routes
		kafkaservice.StartKafkaConsumer()

	    routes.ProductRoutes(router)
    	router.Run(":" + port)

}