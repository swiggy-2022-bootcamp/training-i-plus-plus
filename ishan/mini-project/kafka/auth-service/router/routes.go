package router

import (
	"swiggy/gin/services/auth"

	"github.com/gin-gonic/gin"
)

func ApplyRoutes() *gin.Engine {
	router := gin.Default()
	auth.AuthRoutes((router))
	return router
}
