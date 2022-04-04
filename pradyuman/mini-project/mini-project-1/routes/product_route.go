package routes

import ( 
	"mini-project/controllers"
	"github.com/gin-gonic/gin"
)
func ProductRoute(router *gin.Engine)  {
	router.POST("/product", controllers.AddProduct())
	router.GET("/product/:productId", controllers.GetAProduct())
	router.PUT("/product/:productId", controllers.EditAProduct())
	router.DELETE("/product/:productId", controllers.DeleteAProduct())
	router.GET("/products", controllers.GetAllProducts())
}