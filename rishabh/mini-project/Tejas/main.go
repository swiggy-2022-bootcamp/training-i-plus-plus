package main

import (
	"os"

	"tejas/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	PORT := os.Getenv("PORT")
	router := gin.Default()
	routes.AuthRoutes(router)

	router.Run(":" + PORT)
}
