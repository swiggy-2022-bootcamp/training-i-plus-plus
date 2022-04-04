# Healthcare App

A mini project build to showcase the skills and technologies learnt across Swiggy i++ programme.

## Features implemented

1. [X] Microservice Architecture and Concurrency: Both microservices are run from the same main.go as seperate coroutines with two channels to manage errors, although this can be seperated into two golang programs if needed.
2. [X] REST API: Basic method of communication across and between microservices.
3. [X] JSON: Medium for transferring data across the project
4. [X] Database (Mongo): The Key-Value database used as the backend database.
5. [ ] Sonarqube
6. [X] Kafka
7. [X] Swagger UI: Implemented using gin-swagger and swag. Config stored in ./docs/swagger.json and ./docs/swagger.yaml

## Screenshots

### Gin and Logs

| ![Gin Logs of different Levels](https://raw.githubusercontent.com/swiggy-2022-bootcamp/training-i-plus-plus/main/dhiman/mini-project/screenshots/logs.png) |
| --- |

### Mongo DB

| ![Mongo DB Collection](https://raw.githubusercontent.com/swiggy-2022-bootcamp/training-i-plus-plus/main/dhiman/mini-project/screenshots/mongo_collection.png) |
| --- |

### Kafka

| [![Users Microservice Kafka Producer](https://raw.githubusercontent.com/swiggy-2022-bootcamp/training-i-plus-plus/main/dhiman/mini-project/screenshots/kafka_producer.png)](https://raw.githubusercontent.com/swiggy-2022-bootcamp/training-i-plus-plus/main/dhiman/mini-project/screenshots/kafka_producer.png) | [![Bookkeeping Microservice Kafka Consumer](https://raw.githubusercontent.com/swiggy-2022-bootcamp/training-i-plus-plus/main/dhiman/mini-project/screenshots/kafka_consumer.png)](https://raw.githubusercontent.com/swiggy-2022-bootcamp/training-i-plus-plus/main/dhiman/mini-project/screenshots/kafka_consumer.png) |
| --- | --- |

### Swagger UI

| [![Bookkeeping Microservice Swagger](https://raw.githubusercontent.com/swiggy-2022-bootcamp/training-i-plus-plus/main/dhiman/mini-project/screenshots/bookkeeping_swagger.png)](https://raw.githubusercontent.com/swiggy-2022-bootcamp/training-i-plus-plus/main/dhiman/mini-project/screenshots/bookkeeping_swagger.png) | [![Users Microservice Swagger](https://raw.githubusercontent.com/swiggy-2022-bootcamp/training-i-plus-plus/main/dhiman/mini-project/screenshots/users_swagger.png)](https://raw.githubusercontent.com/swiggy-2022-bootcamp/training-i-plus-plus/main/dhiman/mini-project/screenshots/users_swagger.png) |
| --- | --- |
