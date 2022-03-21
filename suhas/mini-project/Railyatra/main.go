package main

import (
	"gin-mongo-api/config" //add this
	"gin-mongo-api/routes"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	//routes
	routes.AuthRoute(router)
	routes.UserRoute(router)  //add this
	routes.AdminRoute(router) //add this
	return router
}

func main() {

	rout := SetupRouter()
	//run database
	config.ConnectDB()

	rout.Run("localhost:6000")
}
