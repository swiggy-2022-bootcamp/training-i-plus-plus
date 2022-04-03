package configs

import (
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var envPath string = "/Users/dhimanseal/projects/training-i-plus-plus/dhiman/mini-project/.env"

func EnvMongoURI() string {
    err := godotenv.Load(envPath)
    if err != nil {
        log.Fatal("Error loading .env file")
    }
  
    return os.Getenv("MONGOURI")
}

// Get the name of the Clients Collection in MongoDB
func ClientsCollectionName() string {
    err := godotenv.Load(envPath)
    if err != nil {
        log.Fatal("Error loading .env file")
    }
  
    return os.Getenv("CLIENTS_COLLECTION")
}

// Get the name of the Experts Collection in MongoDB
func ExpertsCollectionName() string {
    err := godotenv.Load(envPath)
    if err != nil {
        log.Fatal("Error loading .env file")
    }
  
    return os.Getenv("EXPERTS_COLLECTION")
}
