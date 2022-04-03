package main

import (
	"io"
	
	
	"log"
	routes "bhatiachahat/track-stream-service/routes"
	utils "bhatiachahat/track-stream-service/utils"
   kafkaservice "bhatiachahat/track-stream-service/kafkaservice"
    "os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "bhatiachahat/track-stream-service/docs"




)
// @title          Online Shopping Application -Track Stream Module
// @version        1.0
// @description    This microservice is for track stream module in the online shopping application.
// @contact.name   Chahat Bhatia
// @contact.email  chahatbhatia2014@gmail.com
// @license.name   Apache 2.0
// @host           localhost:8084
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
		port="8084"}

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
	
		docs.SwaggerInfo.Title = "Online Shopping App - Track Stream Module"
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	
		kafkaservice.StartKafkaConsumer()

	    //routes
	    routes.TrackStreamRoutes(router)
    	router.Run(":" + port)

}