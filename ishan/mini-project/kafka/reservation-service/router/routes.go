package router

import (
	"swiggy/gin/services/auth"
	"swiggy/gin/services/reservation"

	"github.com/gin-gonic/gin"
)

func ApplyRoutes() *gin.Engine {
	router := gin.Default()
	auth.AuthRoutes((router))
	reservation.ReservationRoutes(router)
	return router
}
