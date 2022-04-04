package routes

import ( 
	"inventoryService/controllers"
	"inventoryService/middleware"
	"github.com/gin-gonic/gin"
)

func ProductViewRoute(router *gin.Engine)  {
	router.Use(middleware.IsUserAuthorized([]string{"BUYER","SELLER"}))
	
	router.GET("/products", controllers.GetAllProducts())
	router.GET("/product/:productId", controllers.GetAProduct())
	
}

func ProductUpdateRoute(router *gin.Engine)  {
	router.Use(middleware.IsUserAuthorized([]string{"SELLER"}))
	
	router.POST("/product", controllers.AddProduct())
	router.PUT("/product/:productId", controllers.EditAProduct())
	router.DELETE("/product/:productId", controllers.DeleteAProduct())	

}