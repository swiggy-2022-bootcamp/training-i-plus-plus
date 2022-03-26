package main

import (
	"srctc/config"
	"srctc/routes"

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
	userroute.Run("localhost:6000")

	adminroute := SetupAdminRouter()
	adminroute.Run("localhost:6001")

	authroute := SetupAuthRouter()
	authroute.Run("localhost:6002")

}
