package routes

import(
	usercontroller "github.com/bhatiachahat/mongoapi/controllers/users"
	productcontroller "github.com/bhatiachahat/mongoapi/controllers/products"


	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("users/signup", usercontroller.Signup())
	incomingRoutes.POST("users/login", usercontroller.Login())
	 incomingRoutes.GET("/products", productcontroller.GetAllProducts())
	 incomingRoutes.GET("products/:product_id", productcontroller.GetProduct())
}