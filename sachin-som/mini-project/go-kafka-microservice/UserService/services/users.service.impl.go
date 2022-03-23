package services

import (
	"context"

	"github.com/go-kafka-microservice/UserService/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	Ctx            context.Context
	UserCollection *mongo.Collection
}

func NewUserServiceImpl(userCollection *mongo.Collection, ctx context.Context) *UserServiceImpl {
	return &UserServiceImpl{
		Ctx:            ctx,
		UserCollection: userCollection,
	}
}

func (us *UserServiceImpl) CreateUser(user *models.User) error {
	return nil
}

func (us *UserServiceImpl) GetUser(userId int) (*models.User, error) {
	return nil, nil
}
