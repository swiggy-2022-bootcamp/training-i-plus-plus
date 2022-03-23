package services

import (
	"github.com/go-kafka-microservice/UserService/models"
)

/*
* User functions definitions.
 */
type UserService interface {
	CreateUser(*models.User) error
	GetUser(int) (*models.User, error)
}
