package main

import (
	"context"

	docs "github.com/dhi13man/healthcare-app/bookkeeping_service/docs"
	bookkeeping_routes "github.com/dhi13man/healthcare-app/bookkeeping_service/routes"
	"github.com/dhi13man/healthcare-app/bookkeeping_service/services"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func main() {
	docs.SwaggerInfo.BasePath = bookkeeping_routes.BaseURL
	// Set up routes for Bookkeeping Microservice
	bookkeepingRouter := gin.Default()
	bookkeeping_routes.GenerateBookKeepingServiceRoutes(bookkeepingRouter)
	bookkeepingRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Set up Kafka listener
	ctx := context.Background()
	go services.Consume("diagnosis", services.DeserializeAndSaveDiseaseDiagnosis, ctx)
	
	// Run Microservice
	bookkeepingRouter.Run("localhost:8082")
}
