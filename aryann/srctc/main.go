package main

import (
	"os"

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

	router.Run(":" + PORT)
}
