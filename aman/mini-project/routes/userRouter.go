package routes

import (
	"aman-swiggy-mini-project/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	r.GET("/users/:id", controllers.GetUser())
	r.POST("/users/signup", controllers.SignUp())
	r.POST("/users/login", controllers.Login())
	r.POST("/users/logout", controllers.Logout())
}
