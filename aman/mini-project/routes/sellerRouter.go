package routes

import (
	"aman-swiggy-mini-project/controllers"

	"github.com/gin-gonic/gin"
)

func SellerRoutes(r *gin.Engine) {
	r.GET("/sellers/:id", controllers.GetSeller())
	r.POST("/sellers/signup", controllers.SellerSignUp())
	r.POST("/sellers/login", controllers.SellerLogin())
}
