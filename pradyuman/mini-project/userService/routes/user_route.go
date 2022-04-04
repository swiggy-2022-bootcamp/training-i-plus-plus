package routes

import ( 
	"userService/controllers"
	"userService/middleware"
	"github.com/gin-gonic/gin"

)
func UserRoute(router *gin.Engine)  {
	//router.POST("/user", controllers.AddUser())
	router.Use(middleware.IsUserAuthorized([]string{"BUYER","SELLER"}))
	router.GET("/user/:userId", (controllers.GetAUser()))
	router.PUT("/user/:userId", controllers.EditAUser())
	router.DELETE("/user/:userId", controllers.DeleteAUser())
	router.GET("/users", controllers.GetAllUsers())
}