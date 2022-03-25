package main

import (
	"swiggy/gin/services/auth"
	"swiggy/gin/services/reservation"
	"swiggy/gin/services/train"

	"github.com/gin-gonic/gin"
)

func ApplyRoutes(router *gin.Engine) {
	auth.AuthRoutes(router)
	train.TrainRoutes(router)
	reservation.ReservationRoutes(router)
}
