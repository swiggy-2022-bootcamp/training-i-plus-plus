package main

import (
	"time"

	bookkeeping_routes "github.com/dhi13man/healthcare-app/bookkeeping_service/routes"
	user_routes "github.com/dhi13man/healthcare-app/users_service/routes"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
    swaggerFiles "github.com/swaggo/files" // swagger embed files
	docs "github.com/dhi13man/healthcare-app/docs"
)

func main() {
	docs.SwaggerInfo.BasePath = "/"
	// Set up routes for both microsevices
	// General
	router := gin.Default()
	generateBaseRoute(router)
	// Users Microservice
	usersRouter := gin.Default()
	user_routes.GenerateUsersServiceRoutes(usersRouter)
	usersRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Bookkeeping Microservice
	bookkeepingRouter := gin.Default()
	bookkeeping_routes.GenerateBookKeepingServiceRoutes(bookkeepingRouter)
	bookkeepingRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Make channels to get error
	errChanUsers := make(chan error)
	errChanBookKeeping := make(chan error)
	
	// Run Microservices
	go func() {
		router.Run("localhost:8080")
	}()
	go func() {
		errChanUsers <- usersRouter.Run("localhost:8081")
	}()
	go func() {
		errChanBookKeeping <- bookkeepingRouter.Run("localhost:8082")
	}()

	// Listen to errors.
	select {
		case err := <-errChanUsers: // Received in Channel 1
			log.Fatal("users Microservice failed", err)
		case err := <-errChanBookKeeping: // Received in channel 2
			log.Fatal("bookkeeping Microservice failed", err)
		default:
			// Block main thread for this time so goroutines can run with their seperate microservices.
			select {}

	}
}

// Base route response telling the user about available microservices and their paths.
func generateBaseRoute(router *gin.Engine) {
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome to Healthcare App. Below are the two Microservices developed.",
			"routes": []string{
				"/bookkeeping/",
				"/users/",
			},
		})
	})
}
