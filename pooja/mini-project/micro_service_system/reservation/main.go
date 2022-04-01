package main

import (
	"reservation/database"
	"reservation/routes"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	router := gin.Default()
	database.DatabaseConn()

	log.Info("Reservation module is connected to db")

	routes.ReservationRoutes(router)
	router.Run(":6003")
}
