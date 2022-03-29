package router

import (
	"swiggy/gin/services/train"

	"github.com/gin-gonic/gin"
)

func ApplyRoutes() *gin.Engine {
	router := gin.Default()
	train.TrainRoutes((router))
	return router
}
