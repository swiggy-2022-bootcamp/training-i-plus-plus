package main

import (
	"TrainService/database"
	"TrainService/logger"
	"TrainService/routes"

	"github.com/gin-gonic/gin"
)

var logger1 = logger.NewLoggerService("main")

func TrainRouter() *gin.Engine {
	router := gin.Default()
	routes.TrainRoute(router)
	return router
}

func main() {

	database.ConnectDB()
	logger1.Log("Connected to MongoDB")

	trainroute := TrainRouter()
	trainroute.Run("localhost:7000")
}
