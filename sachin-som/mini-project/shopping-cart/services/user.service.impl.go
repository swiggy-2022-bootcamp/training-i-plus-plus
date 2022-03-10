package services

import (
	"context"

	"github.com/sachinsom93/shopping-cart/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImplementation struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(userCollection *mongo.Collection, ctx context.Context) *UserServiceImplementation {
	return &UserServiceImplementation{
		userCollection: userCollection,
		ctx:            ctx,
	}
}

func (u *UserServiceImplementation) CreateUser(user *models.User) error {
	return nil
}

func (u *UserServiceImplementation) GetUser(email *string) (*models.User, error) {
	return nil, nil
}

func (u *UserServiceImplementation) GetAllUser() []*models.User {
	return nil
}

func (u *UserServiceImplementation) UpdateUser(user *models.User) error {
	return nil
}

func (u *UserServiceImplementation) DeleteUser(email *string) error {
	return nil
}
