package main

import (
	"log"
	"os"
	"userService/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failed to load .env file", err.Error())
	}
	PORT := os.Getenv("PORT")
	router := gin.Default()

	routes.UserRouter(router)

	router.Run(":" + PORT)
}
