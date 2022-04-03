package config

import (
	log "authService/logger"
	"os"

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
