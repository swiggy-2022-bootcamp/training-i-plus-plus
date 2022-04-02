package main

import (
	"UserService/database"
	"UserService/logger"
	"UserService/routes"

	"github.com/gin-gonic/gin"
)

var logger1 = logger.NewLoggerService("main")

func UserRouter() *gin.Engine {
	router := gin.Default()
	routes.UserRoute(router)
	return router
}

func main() {

	database.ConnectDB()
	logger1.Log("Connected to MongoDB")

	userroute := UserRouter()
	userroute.Run("localhost:6000")
}
