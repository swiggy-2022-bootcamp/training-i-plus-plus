package router

import (
	"swiggy/gin/services/auth"
	"swiggy/gin/services/reservation"
	"swiggy/gin/services/train"

	"github.com/gin-gonic/gin"
)

func ApplyRoutes() *gin.Engine {
	router := gin.Default()

	auth.AuthRoutes(router)
	train.TrainRoutes(router)
	reservation.ReservationRoutes(router)
	return router
}
