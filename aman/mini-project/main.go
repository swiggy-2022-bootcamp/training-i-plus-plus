package main

import (
	_ "aman-swiggy-mini-project/docs"
	"aman-swiggy-mini-project/models"
	"aman-swiggy-mini-project/routes"
	"encoding/json"
	"fmt"
	"io/ioutil"

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
	r.POST("/", test)
	r.Run()
}

func test(c *gin.Context) {
	body := c.Request.Body
	decoder := json.NewDecoder(body)
	var user1 models.User
	err := decoder.Decode(&user1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user1)
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, gin.H{
		"HI":   value,
		"body": user1,
	})
}
