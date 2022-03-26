package main

import (
	"srctc/config"
	"srctc/routes"

	"github.com/gin-gonic/gin"
)

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

	config.ConnectDB()

	userroute := UserRouter()
	userroute.Run("localhost:8000")

	adminroute := AdminRouter()
	adminroute.Run("localhost:8001")

	authroute := AuthRouter()
	authroute.Run("localhost:8002")

}
