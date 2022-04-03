package main

import (
	"adminService/config"
	log "adminService/logger"
	"fmt"

	"adminService/controllers"
	"adminService/routes"

	"github.com/gin-gonic/gin"
)

var (
	errLog = log.InfoLogger.Println
)

func main() {
	config.ConnectDB()

	go func() {
		err := controllers.StartAdmingrpc()
		if err != nil {
			fmt.Println(err)
			return
		}
	}()

	router := gin.Default()
	routes.AdminRoute(router)
	router.Run("localhost:6001")
}
