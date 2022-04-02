package main

import (
	"TicketService/database"
	"TicketService/logger"
	"TicketService/routes"

	"github.com/gin-gonic/gin"
)

var logger1 = logger.NewLoggerService("main")

func TicketRouter() *gin.Engine {
	router := gin.Default()
	routes.TicketRoute(router)
	return router
}

func main() {

	database.ConnectDB()
	logger1.Log("Connected to MongoDB")

	userroute := TicketRouter()
	userroute.Run("localhost:8000")
}
