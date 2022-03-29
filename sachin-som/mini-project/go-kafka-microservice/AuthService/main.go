package main

import (
	"fmt"
	"log"

	"github.com/go-kafka-microservice/AuthService/services"
	"github.com/go-kafka-microservice/AuthService/utils"
	"github.com/joho/godotenv"
)

var (
	err          error
	jwtUtils     utils.JWTUtils
	authServices services.AuthServices
)

func init() {
	// Parse .dot file env variables
	if err = godotenv.Load(); err != nil {
		log.Fatal("Error Loading in .env file: ", err.Error())
	}
	// Initialize jwt utils service
	jwtUtils = utils.NewJWTUtils()

	// Initialize authService
	authServices = services.NewAuthServiceImpl(jwtUtils)
}
func main() {
	fmt.Println("Auth Service.")
}
