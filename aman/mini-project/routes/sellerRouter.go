package routes

import (
	"aman-swiggy-mini-project/controllers"

	"github.com/gin-gonic/gin"
)

func SellerRoutes(r *gin.Engine) {
	r.GET("/sellers/:id", controllers.GetUser())
	r.POST("/sellers/signup", controllers.SignUp())
	r.POST("/sellers/login", controllers.Login())
}
