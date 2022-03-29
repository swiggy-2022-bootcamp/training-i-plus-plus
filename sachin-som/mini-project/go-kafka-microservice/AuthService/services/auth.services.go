package services

import "github.com/go-kafka-microservice/AuthService/models"

type AuthServices interface {
	Authenticate(*models.Credentials) (string, error)
	Authorize(string) error
}
