package routes

import(
	productcontroller "github.com/bhatiachahat/mongoapi/controllers/products"
	 "github.com/bhatiachahat/mongoapi/middleware"

	
	"github.com/gin-gonic/gin"
)

func ProductRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate())
	 incomingRoutes.POST("/products/add", productcontroller.AddProduct())
     incomingRoutes.DELETE("/products/:product_id", productcontroller.DeleteProduct())
	 incomingRoutes.PUT("/products/:product_id", productcontroller.UpdateProduct())
}