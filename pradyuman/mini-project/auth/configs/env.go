package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

//checks if the environment variable is correctly loaded and returns the environment variable.

func EnvSecretKeyJWT() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("SecretKeyJWT")
}

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("MONGOURI")
}

func EnvPort() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("PORT")
}
