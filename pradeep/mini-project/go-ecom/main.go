package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pswaldia/go-ecom/controllers"
	"github.com/pswaldia1/go-ecom/database"
	"github.com/pswaldia1/go-ecom/middleware"
	"github.com/pswaldia1/go-ecom/routes"
)

func main(){
	port := os.Getenv('PORT')
	if port == ""{
		port = "8080"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))
	
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeItem", app.RemoveItem())
	router.GET("/cartCheckout", app.Checkout())
	router.GET("/instantBuy", app.instantBuy())
	log.Fatal(router.Run(":" + port))
}