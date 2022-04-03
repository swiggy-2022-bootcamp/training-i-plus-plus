package main

import (
	"userService/config"
	"userService/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	router := gin.Default()
	routes.UserRoute(router)
	router.Run("localhost:6002")
}
