package main

import (
	"gin-mongo-api/config" //add this
	"gin-mongo-api/routes"

	"github.com/gin-gonic/gin"
)

func SetupUserRouter() *gin.Engine {
	router := gin.Default()
	routes.UserRoute(router)
	return router
}

func SetupAdminRouter() *gin.Engine {
	router := gin.Default()
	routes.AdminRoute(router)
	return router
}

func SetupAuthRouter() *gin.Engine {
	router := gin.Default()
	routes.AuthRoute(router)
	return router
}

func main() {

	config.ConnectDB()

	userroute := SetupUserRouter()
	go userroute.Run("localhost:6000")

	adminroute := SetupAdminRouter()
	go adminroute.Run("localhost:6001")

	authroute := SetupAuthRouter()
	authroute.Run("localhost:6002")

}
