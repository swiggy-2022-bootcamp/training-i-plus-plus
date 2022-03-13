package routes

import(
	usercontroller "github.com/bhatiachahat/mongoapi/controllers/users"
	cartcontroller "github.com/bhatiachahat/mongoapi/controllers/cart"
	 "github.com/bhatiachahat/mongoapi/middleware"

	
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", usercontroller.GetUsers())
	incomingRoutes.GET("/users/:user_id", usercontroller.GetUser())
	incomingRoutes.GET("/cart/all/:user_id",cartcontroller.GetAllCartItems())
	incomingRoutes.DELETE("/cart/:cart_id",cartcontroller.DeleteFromCart());
	incomingRoutes.POST("/cart/add",cartcontroller.AddProductToCart())

	
}