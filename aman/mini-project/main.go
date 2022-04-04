package main

import (
	_ "aman-swiggy-mini-project/docs"
	"aman-swiggy-mini-project/routes"
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Swagger service APIs for Shopping App
// @version 1.0
// @description Swagger APIs available for Shopping App.
// @termsOfService http://github.com/justamangupta

// @contact.name API Support
// @contact.url http://github.com/justamangupta

// @host localhost:8080
// @BasePath /

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name token
func main() {
	fmt.Println("Starting Project")
	r := gin.New()
	r.Use(gin.Logger())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.UserRoutes(r)
	routes.SellerRoutes(r)
	routes.ProductRoutes(r)
	routes.CartRoutes(r)
	routes.InventoryRoutes(r)
	routes.CartItemsRoutes(r)
	routes.OrderRoutes(r)
	routes.RequestRoutes(r)
	r.Run()
}
