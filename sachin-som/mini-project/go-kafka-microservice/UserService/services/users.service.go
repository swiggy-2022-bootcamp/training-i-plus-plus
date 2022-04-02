package services

import (
	"github.com/go-kafka-microservice/UserService/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
* User functions definitions.
 */
type UserService interface {
	CreateUser(*models.User) (string, error)
	GetUser(primitive.ObjectID) (*models.User, error)
	Login(*models.Credentials) (string, error)
	UpdateUser(primitive.ObjectID, *models.User) error
	DeleteUser(primitive.ObjectID) error
}
