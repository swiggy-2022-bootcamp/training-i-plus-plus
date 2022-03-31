package configs

import (
    "log"
    "os"
    "github.com/joho/godotenv"
)

var envPath string = "/Users/dhimanseal/projects/training-i-plus-plus/dhiman/mini-project/.env"

func EnvMongoURI() string {
    err := godotenv.Load(envPath)
    if err != nil {
        log.Fatal("Error loading .env file")
    }
  
    return os.Getenv("MONGOURI")
}

// Get the name of the Users Collection in MongoDB
func UsersCollectionName() string {
    err := godotenv.Load(envPath)
    if err != nil {
        log.Fatal("Error loading .env file")
    }
  
    return os.Getenv("USERS_COLLECTION")
}

// Get the name of the Medicines Collection in MongoDB
func MedicinesCollectionName() string {
    err := godotenv.Load(envPath)
    if err != nil {
        log.Fatal("Error loading .env file")
    }
  
    return os.Getenv("MEDICINES_COLLECTION")
}