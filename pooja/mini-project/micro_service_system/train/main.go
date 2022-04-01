package main

import (
	"train/database"
	"train/routes"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	router := gin.Default()
	database.DatabaseConn()

	log.Info("Train module is connected to db")

	routes.TrainRoutes(router)
	router.Run(":6002")
}
