package configs

import (
	"os"

	"go.uber.org/zap"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
    err := godotenv.Load()
    if err != nil {
        zap.L().Error("Error loading .env file")
    }
  
    return os.Getenv("MONGOURI")
}