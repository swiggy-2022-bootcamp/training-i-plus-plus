package services

import (
	"context"

	"github.com/go-kafka-microservice/UserService/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	user.ID = primitive.NewObjectID()
	if _, err := us.UserCollection.InsertOne(us.Ctx, user); err != nil {
		return err
	}
	return nil
}

func (us *UserServiceImpl) GetUser(userId primitive.ObjectID) (*models.User, error) {
	filter := bson.D{bson.E{Key: "_id", Value: userId}}
	var user models.User
	if err := us.UserCollection.FindOne(us.Ctx, filter).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
