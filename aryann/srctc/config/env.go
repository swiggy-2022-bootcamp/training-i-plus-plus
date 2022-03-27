package config

import (
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
		// log.Fatal("Error loading .env file")
	}

	return os.Getenv("MONGOURI")
}
