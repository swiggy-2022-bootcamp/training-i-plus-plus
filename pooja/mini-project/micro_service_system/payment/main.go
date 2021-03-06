package payment

import (
	"payment/database"
	"payment/routes"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	router := gin.Default()

	database.DatabaseConn()
	log.Info("Connected to Payment db")
	routes.PaymentRoutes(router)
	router.Run(":6004")
}
