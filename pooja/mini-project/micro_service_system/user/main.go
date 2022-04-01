package main

import (
	"user/database"
	"user/routes"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	router := gin.Default()
	database.DatabaseConn()

	log.Info("Connected to db")

	routes.UserRoutes(router)
	router.Run(":6001")
}
