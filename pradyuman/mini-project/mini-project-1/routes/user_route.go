package routes

import ( 
	"mini-project/controllers"
	"github.com/gin-gonic/gin"
)
func UserRoute(router *gin.Engine)  {
	//router.POST("/user", controllers.AddUser())
	router.Use(controllers.IsAuthorized())
	router.GET("/user/:userId", (controllers.GetAUser()))
	router.PUT("/user/:userId", controllers.EditAUser())
	router.DELETE("/user/:userId", controllers.DeleteAUser())
	router.GET("/users", controllers.GetAllUsers())
}