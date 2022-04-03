package configs

import (
    "log"
    "os"
    "github.com/joho/godotenv"
)

// Path where the environment variable is located.
var envPath string = "/Users/dhimanseal/projects/training-i-plus-plus/dhiman/mini-project/.env"

// Get the URi of the Mongo Database.
func EnvMongoURI() string {
    err := godotenv.Load(envPath)
    if err != nil {
        log.Fatal("Error loading .env file.")
    }
  
    return os.Getenv("MONGOURI")
}

// Get the name of the Users Collection in MongoDB
func DiseasesCollectionName() string {
    err := godotenv.Load(envPath)
    if err != nil {
        log.Fatal("Error loading .env file")
    }
  
    return os.Getenv("DISEASES_COLLECTION")
}

// Get the name of the Medicines Collection in MongoDB
func MedicinesCollectionName() string {
    err := godotenv.Load(envPath)
    if err != nil {
        log.Fatal("Error loading .env file")
    }
  
    return os.Getenv("MEDICINES_COLLECTION")
}