package main

import (
	"fmt"
	"srctc/config"
	"srctc/routes"

	"github.com/gin-gonic/gin"
)

func UserRouter() *gin.Engine {
	fmt.Println("user check")
	router := gin.Default()
	routes.UserRoute(router)
	return router
}

func AdminRouter() *gin.Engine {
	fmt.Println("admin check")
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
	userroute.Run("localhost:6000")

	adminroute := AdminRouter()
	adminroute.Run("localhost:6001")

	authroute := AuthRouter()
	authroute.Run("localhost:6002")

}
