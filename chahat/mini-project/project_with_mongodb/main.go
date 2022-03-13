package main

import (
	
	
	"log"
    "github.com/bhatiachahat/mongoapi/routes"
    "os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)





func main() {
	

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	if port==""{
		port="8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)
	routes.ProductRoutes(router)

	

	router.Run(":" + port)

}