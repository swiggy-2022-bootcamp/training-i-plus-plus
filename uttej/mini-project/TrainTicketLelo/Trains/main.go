package main

import (
	"Trains/config"
	controllers "Trains/controller"
	"Trains/docs"
	"Trains/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Swagger TrainTicketLelo Trains Service
// @version         1.0
// @description     Swagger TrainTicketLelo Trains Service
// @termsOfService  http://swagger.io/terms/

// @contact.name   Uttej Immadi
// @contact.url    http://www.swagger.io/support
// @contact.email  immadiuttej@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8003

// @securityDefinitions.apiKey ApiKeyAuth
// @type apiKey
// @in header
// @name Authorization
func main() {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/trains", middleware.IfAuthorized(controllers.CreateTrain))
	r.GET("/trains", middleware.IfAuthorized(controllers.GetTrains))
	r.GET("/trains/:trainId", middleware.IfAuthorized(controllers.GetTrainById))
	r.PUT("/trains/:trainId", middleware.IfAuthorized(controllers.UpdateTrainById))
	r.DELETE("/trains/:trainId", middleware.IfAuthorized(controllers.DeleteTrainbyId))
	r.POST("/trains/:trainId/:updateCount", middleware.IfAuthorized(controllers.UpdateTicketCount))
	docs.SwaggerInfo.Title = "Swagger TrainTicketLelo Trains Service"

	portAddress := ":" + config.TrainServicePort
	r.Run(portAddress)

}
