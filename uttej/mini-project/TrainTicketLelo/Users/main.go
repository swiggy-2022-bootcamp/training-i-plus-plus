package main

import (
	"Users/config"
	controller "Users/controllers"
	"Users/docs"
	"Users/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/signup", controller.CreateUser)

	r.POST("/users/login", controller.LogInUser)
	r.GET("/users", middleware.IfAuthorized(controller.GetAllUsers))
	r.GET("/users/:userId", middleware.IfAuthorized(controller.GetUserById))
	r.PUT("/users/:userId", middleware.IfAuthorized(controller.UpdateUserById))
	r.DELETE("/users/:userId", middleware.IfAuthorized(controller.DeleteUserbyId))
	docs.SwaggerInfo.Title = "Swagger TrainTicketLelo Users Service"

	portAddress := ":" + config.UserServicePort
	r.Run(portAddress)
}

// @title           Swagger TrainTicketLelo Users Service
// @version         1.0
// @description     Swagger TrainTicketLelo Users Service
// @termsOfService  http://swagger.io/terms/

// @contact.name   Uttej Immadi
// @contact.url    http://www.swagger.io/support
// @contact.email  immadiuttej@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8001

// @securityDefinitions.apiKey ApiKeyAuth
// @type apiKey
// @in header
// @name Authorization
