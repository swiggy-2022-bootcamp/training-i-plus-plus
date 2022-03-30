package main

import (
	"log"
	"os"
	"userService/routes"

	"userService/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title           Swagger Train-Ticket Booking System API
// @version         1.0
// @description     Swagger Train-Ticket Booking System API.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  swiggyb2013@datascience.manipal.edu

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8081

// @securityDefinitions.basic  BasicAuth
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failed to load .env file", err.Error())
	}

	docs.SwaggerInfo.Title = "Swagger Train-Ticket Booking System API"

	PORT := os.Getenv("PORT")
	router := gin.Default()

	routes.UserRouter(router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":" + PORT)
}
