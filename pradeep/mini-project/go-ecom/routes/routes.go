package routes

import (
	"github.com/gin-gonic/gin"
	"githun.com/pswaldia/go-ecom/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.Signup())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("admin/addproduct", controllers.AddProduct())
	incomingRoutes.GET("/users/productview", controllers.ProductView())
	incomingRoutes.GET("/users/search", controllers.SearchProduct())
}
