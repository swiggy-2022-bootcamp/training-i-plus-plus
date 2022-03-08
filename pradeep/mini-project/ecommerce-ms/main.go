package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pswaldia/ecommerce-ms/controllers"
	"github.com/pswaldia/ecommerce-ms/database"
	"github.com/pswaldia/ecommerce-ms/middleware"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))
	router := gin.New()

	router.Use(gin.Logger())
	router.UseRoutes(router)
	router.Use(middleware.Authentication())

}
