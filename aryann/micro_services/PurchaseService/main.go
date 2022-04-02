package main

import (
	"PurchaseService/database"
	"PurchaseService/logger"
	"PurchaseService/routes"

	"github.com/gin-gonic/gin"
)

var logger1 = logger.NewLoggerService("main")

func PurchaseRouter() *gin.Engine {
	router := gin.Default()
	routes.PurchaseRoute(router)
	return router
}

func main() {

	database.ConnectDB()
	logger1.Log("Connected to MongoDB")

	userroute := PurchaseRouter()
	userroute.Run("localhost:6000")

}
