package main

import (
	docs "github.com/dhi13man/healthcare-app/users_service/docs"
	user_routes "github.com/dhi13man/healthcare-app/users_service/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func main() {
	docs.SwaggerInfo.BasePath = "/"
	// Set up routes for both microsevices
	// Users Microservice
	usersRouter := gin.Default()
	user_routes.GenerateUsersServiceRoutes(usersRouter)
	usersRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Run Microservice
	usersRouter.Run("localhost:8081")
}
