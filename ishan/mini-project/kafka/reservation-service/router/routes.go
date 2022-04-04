package router

import (
	"swiggy/gin/services/reservation"

	"github.com/gin-gonic/gin"
)

func ApplyRoutes() *gin.Engine {
	router := gin.Default()
	reservation.ReservationRoutes(router)
	return router
}
