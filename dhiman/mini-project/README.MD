# Healthcare App

A mini project build to showcase the skills and technologies learnt across Swiggy i++ programme.

## Features implemented

1. [X] Microservice Architecture and Concurrency: Both microservices are run from the same main.go as seperate coroutines with two channels to manage errors, although this can be seperated into two golang programs if needed.
2. [X] REST API: Basic method of communication across and between microservices.
3. [X] JSON: Medium for transferring data across the project
4. [X] Database (Mongo): The Key-Value database used as the backend database.
5. [X] Kafka: To send messages from users microservice to bookkeeping microservice.
6. [X] Swagger UI: Implemented using gin-swagger and swag. Config stored in ./docs/swagger.json and ./docs/swagger.yaml
7. [X] Sonarqube: For Code coverage and code quality.

## Usage

1. Set the environment variables in a .env file:
   1. **MongoDB:** `MONGOURI`, `DISEASES_COLLECTION`, `EXPERTS_COLLECTION`, `MEDICINES_COLLECTION`, `CLIENTS_COLLECTION`
   2. **Kafka:** `KAFKA_BROKER_ADDRESS`, `KAFKA_DIAGNOSIS_TOPIC`
   3. **Microservice Addresses:** `USERS_SERVICE_ADDRESS`, `BOOKKEEPING_SERVICE_ADDRESS`

### Starting as Microservices

1. **Run Users Microservice:** `go run "./users_service/main.go"`
2. **Run Bookkeeping Microservice:** `go run "./bookkeeping_service/main.go"`

### Starting as a Monolith using Coroutines

1. Delete `go.mod` in `./bookkeeping_service` and `./users_service` so golang can function without the nested modules.
2. **Run base Monolith code:** `go run "./main.go"`

## Screenshots

### Gin and Logging

| Gin and Logging in different Levels using [logrus](https://github.com/sirupsen/logrus) |
| --- |
| ![Gin Logs of different Levels](https://raw.githubusercontent.com/swiggy-2022-bootcamp/training-i-plus-plus/main/dhiman/mini-project/screenshots/logs.png) |

### Mongo DB

| Mongo DB Collections |
| --- |
| ![Mongo DB Collection](https://raw.githubusercontent.com/swiggy-2022-bootcamp/training-i-plus-plus/main/dhiman/mini-project/screenshots/mongo_collection.png) |

### Kafka

| Users Microservice (Kafka Producer) | Bookkeeping Microservice (Kafka Consumer) |
| --- | --- |
| [![Users Microservice Kafka Producer](https://raw.githubusercontent.com/swiggy-2022-bootcamp/training-i-plus-plus/main/dhiman/mini-project/screenshots/kafka_producer.png)](https://raw.githubusercontent.com/swiggy-2022-bootcamp/training-i-plus-plus/main/dhiman/mini-project/screenshots/kafka_producer.png) | [![Bookkeeping Microservice Kafka Consumer](https://raw.githubusercontent.com/swiggy-2022-bootcamp/training-i-plus-plus/main/dhiman/mini-project/screenshots/kafka_consumer.png)](https://raw.githubusercontent.com/swiggy-2022-bootcamp/training-i-plus-plus/main/dhiman/mini-project/screenshots/kafka_consumer.png) |

### Swagger UI

| Bookkeeping Microservice Swagger | Users Microservice Swagger |
| --- | --- |
| [![Bookkeeping Microservice Swagger](https://raw.githubusercontent.com/swiggy-2022-bootcamp/training-i-plus-plus/main/dhiman/mini-project/screenshots/bookkeeping_swagger.png)](https://raw.githubusercontent.com/swiggy-2022-bootcamp/training-i-plus-plus/main/dhiman/mini-project/screenshots/bookkeeping_swagger.png) | [![Users Microservice Swagger](https://raw.githubusercontent.com/swiggy-2022-bootcamp/training-i-plus-plus/main/dhiman/mini-project/screenshots/users_swagger.png)](https://raw.githubusercontent.com/swiggy-2022-bootcamp/training-i-plus-plus/main/dhiman/mini-project/screenshots/users_swagger.png) |

### Sonarqube and Unit Tests

| Sonarqube Projects Dashboard  | Unit Testing in Golang |
| --- | --- |
| [![Sonarqube Projects Dashboard](https://raw.githubusercontent.com/swiggy-2022-bootcamp/training-i-plus-plus/main/dhiman/mini-project/screenshots/sonarqube.png)](https://raw.githubusercontent.com/swiggy-2022-bootcamp/training-i-plus-plus/main/dhiman/mini-project/screenshots/sonarqube.png) | [![Unit Testing Golang Examples](https://raw.githubusercontent.com/swiggy-2022-bootcamp/training-i-plus-plus/main/dhiman/mini-project/screenshots/tests.png)](https://raw.githubusercontent.com/swiggy-2022-bootcamp/training-i-plus-plus/main/dhiman/mini-project/screenshots/tests.png) |
