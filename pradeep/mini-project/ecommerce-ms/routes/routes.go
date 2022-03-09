package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pswaldia/ecommerce-ms/controllers"
)

func userRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/users/addproduct", controllers.AddProduct())
	incomingRoutes.GET("/users/productview", controllers.GetProduct())
	incomingRoutes.GET("/users/search", controllers.Search())

}
