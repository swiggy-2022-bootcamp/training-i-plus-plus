package services

import (
	"github.com/go-kafka-microservice/UserService/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
* User functions definitions.
 */
type UserService interface {
	CreateUser(*models.User) error
	GetUser(primitive.ObjectID) (*models.User, error)
}
