package routes

import ( 
	"auth/controllers"
	"github.com/gin-gonic/gin"
)
func AuthRoute(router *gin.Engine)  {
	router.POST("/signup", controllers.SignUpUser())
	router.POST("/login", controllers.Login())
}