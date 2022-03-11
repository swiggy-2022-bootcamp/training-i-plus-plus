package services

import (
	"context"
	"errors"
	"time"

	"github.com/sachinsom93/shopping-cart/models"
	"go.mongodb.org/mongo-driver/bson"
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
	_, err := u.userCollection.InsertOne(u.ctx, user)
	return err
}

func (u *UserServiceImplementation) GetUser(email *string) (*models.User, error) {
	var user *models.User
	query := bson.D{bson.E{Key: "email", Value: email}}
	err := u.userCollection.FindOne(u.ctx, query).Decode(&user)
	return user, err
}

func (u *UserServiceImplementation) GetAllUser() ([]*models.User, error) {
	var users []*models.User
	cursor, err := u.userCollection.Find(u.ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(u.ctx) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(u.ctx)
	if len(users) == 0 {
		return nil, errors.New("users not fuond.")
	}
	return users, nil
}

func (u *UserServiceImplementation) UpdateUser(user *models.User) error {
	filterQuery := bson.D{bson.E{Key: "email", Value: user.Email}}
	updatedTime, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateQuery := bson.D{
		bson.E{
			Key:   "first_name",
			Value: user.FirstName,
		},
		bson.E{
			Key:   "last_name",
			Value: user.LastName,
		},
		bson.E{
			Key:   "email",
			Value: user.Email,
		},
		bson.E{
			Key:   "phone",
			Value: user.Phone,
		},
		bson.E{
			Key:   "updated_at",
			Value: updatedTime,
		},
	}
	result, _ := u.userCollection.UpdateOne(u.ctx, filterQuery, updateQuery)
	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update.")
	}
	return nil
}

func (u *UserServiceImplementation) DeleteUser(email *string) error {
	filterQuery := bson.D{bson.E{Key: "email", Value: email}}
	result, _ := u.userCollection.DeleteOne(u.ctx, filterQuery)
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for deleting.")
	}
	return nil
}
