package main

import (
	bookkeeping_routes "github.com/dhi13man/healthcare-app/bookkeeping_service/routes"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
    swaggerFiles "github.com/swaggo/files" // swagger embed files
	docs "github.com/dhi13man/healthcare-app/bookkeeping_service/docs"
)

func main() {
	docs.SwaggerInfo.BasePath = bookkeeping_routes.BaseURL
	// Set up routes for Bookkeeping Microservice
	bookkeepingRouter := gin.Default()
	bookkeeping_routes.GenerateBookKeepingServiceRoutes(bookkeepingRouter)
	bookkeepingRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	// Run Microservice
	bookkeepingRouter.Run("localhost:8082")
}
