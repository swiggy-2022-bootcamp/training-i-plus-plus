package routes

import (
	"aman-swiggy-mini-project/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	r.GET("/users", controllers.GetUsers())
	r.GET("/users/:id", controllers.GetUser())
	r.POST("/users/signup", controllers.SignUp("abc"))
	r.POST("/users/login", controllers.Login("def"))
}
