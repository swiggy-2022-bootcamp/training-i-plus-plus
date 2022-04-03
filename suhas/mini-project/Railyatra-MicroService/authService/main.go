package main

import (
	"authService/config"
	log "authService/logger"
	"fmt"

	"authService/controllers"
	"authService/routes"

	"github.com/gin-gonic/gin"
)

var (
	errLog = log.InfoLogger.Println
)

func main() {
	config.ConnectDB()

	go func() {
		err := controllers.Startgrpc()
		if err != nil {
			fmt.Println(err)
			return
		}
	}()

	router := gin.Default()
	routes.AuthRoute(router)
	router.Run("localhost:6003")
}
