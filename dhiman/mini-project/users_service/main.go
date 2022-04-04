package main

import (
	user_routes "github.com/dhi13man/healthcare-app/users_service/routes"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
    swaggerFiles "github.com/swaggo/files" // swagger embed files
	docs "github.com/dhi13man/healthcare-app/users_service/docs"
)

func main() {
	docs.SwaggerInfo.BasePath = user_routes.BaseURL
	// Set up routes for both microsevices
	// Users Microservice
	usersRouter := gin.Default()
	user_routes.GenerateUsersServiceRoutes(usersRouter)
	usersRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	// Run Microservice
	usersRouter.Run("localhost:8081")
}
