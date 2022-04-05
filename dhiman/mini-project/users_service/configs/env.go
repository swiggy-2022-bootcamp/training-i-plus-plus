package configs

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

// Path where .env is located.
const envPath string = "/Users/dhimanseal/projects/training-i-plus-plus/dhiman/mini-project/.env"

// Error Loading .env file message constant
const errLoadingEnv string = "Error loading .env file."

// Get the Address where the Users Microservice is to be ran.
func UsersServiceAddress() string {
    err := godotenv.Load(envPath)
    if err != nil {
        log.Fatal(errLoadingEnv)
    }
  
    return os.Getenv("USERS_SERVICE_ADDRESS")
}

func EnvMongoURI() string {
    err := godotenv.Load(envPath)
    if err != nil {
        log.Fatal(errLoadingEnv)
    }
  
    return os.Getenv("MONGOURI")
}

// Get the name of the Clients Collection in MongoDB
func ClientsCollectionName() string {
    err := godotenv.Load(envPath)
    if err != nil {
        log.Fatal(errLoadingEnv)
    }
  
    return os.Getenv("CLIENTS_COLLECTION")
}

// Get the name of the Experts Collection in MongoDB
func ExpertsCollectionName() string {
    err := godotenv.Load(envPath)
    if err != nil {
        log.Fatal(errLoadingEnv)
    }
  
    return os.Getenv("EXPERTS_COLLECTION")
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

// Get the name of the Kafka Diagnosis Topic
func BCryptHashRounds() int {
    err := godotenv.Load(envPath)
    if err != nil {
        log.Fatal(errLoadingEnv)
    }
  
    rounds, err := strconv.Atoi(os.Getenv("BCRYPT_HASH_ROUNDS"))
    if err != nil {
        return 10
    } else {
        return rounds
    }
}
