package configs

import (
    "log"
    "os"
    "github.com/joho/godotenv"
)

// Path where .env is located.
const envPath string = "/Users/dhimanseal/projects/training-i-plus-plus/dhiman/mini-project/.env"

// Error Loading .env file message constant
const errLoadingEnv string = "Error loading .env file."

// Get the Address where the Bookkeeping Microservice is to be ran.
func BookkeepingServiceAddress() string {
    err := godotenv.Load(envPath)
    if err != nil {
        log.Fatal(errLoadingEnv)
    }
  
    return os.Getenv("BOOKKEEPING_SERVICE_ADDRESS")
}

// Get the URi of the Mongo Database.
func EnvMongoURI() string {
    err := godotenv.Load(envPath)
    if err != nil {
        log.Fatal(errLoadingEnv)
    }
  
    return os.Getenv("MONGOURI")
}

// Get the name of the Users Collection in MongoDB
func DiseasesCollectionName() string {
    err := godotenv.Load(envPath)
    if err != nil {
        log.Fatal(errLoadingEnv)
    }
  
    return os.Getenv("DISEASES_COLLECTION")
}

// Get the name of the Medicines Collection in MongoDB
func MedicinesCollectionName() string {
    err := godotenv.Load(envPath)
    if err != nil {
        log.Fatal(errLoadingEnv)
    }
  
    return os.Getenv("MEDICINES_COLLECTION")
}

// Get the Address of the Kafka Broker
func KafkaBrokerAddress() string {
    err := godotenv.Load(envPath)
    if err != nil {
        log.Fatal(errLoadingEnv)
    }
  
    return os.Getenv("KAFKA_BROKER_ADDRESS")
}

// Get the name of the Kafka Diagnosis Topic
func KafkaDiagnosisTopic() string {
    err := godotenv.Load(envPath)
    if err != nil {
        log.Fatal(errLoadingEnv)
    }
  
    return os.Getenv("KAFKA_DIAGNOSIS_TOPIC")
}
