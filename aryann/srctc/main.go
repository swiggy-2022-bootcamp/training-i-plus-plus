package main

import (
	"fmt"
	"srctc/database"
	"srctc/logger"
	"srctc/routes"

	"github.com/gin-gonic/gin"
)

var logger1 = logger.NewLoggerService("main")

func UserRouter() *gin.Engine {
	router := gin.Default()
	routes.UserRoute(router)
	return router
}

func AdminRouter() *gin.Engine {
	router := gin.Default()
	routes.AdminRoute(router)
	return router
}

func AuthRouter() *gin.Engine {
	router := gin.Default()
	routes.AuthRoute(router)
	return router
}

func main() {

	database.ConnectDB()
	logger1.Log("Connected to MongoDB")

	fmt.Println("Type  1. User  2. Admin  3. Auth: ")
	var choice int
	fmt.Scanln(&choice)

	if choice == 1 {
		userroute := UserRouter()
		userroute.Run("localhost:6000")
	}
	if choice == 2 {
		adminroute := AdminRouter()
		adminroute.Run("localhost:7000")
	}
	if choice == 3 {
		authroute := AuthRouter()
		authroute.Run("localhost:8000")
	}

}
