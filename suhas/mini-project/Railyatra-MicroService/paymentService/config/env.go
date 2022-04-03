package config

import (
	"os"
	log "paymentService/logger"

	"github.com/joho/godotenv"
)

var (
	errLog = log.ErrorLogger.Println
)

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		errLog("Error loading .env file")
	}

	return os.Getenv("MONGOURI")
}

func EnvStripeSecret() string {
	err := godotenv.Load()
	if err != nil {
		errLog("Error loading .env file")
	}
	return os.Getenv("STRIPE_SECRET_KEY")
}
