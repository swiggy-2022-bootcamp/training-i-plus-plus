package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMonogoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file.")
	}
	return os.Getenv("MONGO_URI")
}

func EnvJWTSecretKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file.")
	}

	return os.Getenv("SECRET_KEY")
}

func EnvPORT() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file.")
	}

	return os.Getenv("PORT")
}
