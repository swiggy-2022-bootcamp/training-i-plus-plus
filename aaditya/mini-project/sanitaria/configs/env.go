package configs

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

func EnvMonogoURI() string{
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file.")
	}

	return os.Getenv("MONGOURI")
}

func EnvJWTSecretKey() string{
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file.")
	}

	return os.Getenv("SECRET_KEY")
}