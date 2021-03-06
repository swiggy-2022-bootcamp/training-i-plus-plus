package services

import (
	"github.com/swastiksahoo153/train-reservation-system/models"
	"go.mongodb.org/mongo-driver/mongo"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"errors"
	"fmt"
)

type UserServiceImpl struct {
	usercollection	*mongo.Collection
	ctx 			context.Context
}

func NewUserService (usercollection *mongo.Collection, ctx context.Context) UserService{
	return &UserServiceImpl{
		usercollection: usercollection,
		ctx:			ctx,
	}
}

func (u *UserServiceImpl) CreateUser(user *models.User) error{
	_, err := u.usercollection.InsertOne(u.ctx, user)
	return err
}

func (u *UserServiceImpl) GetUser(username *string) (*models.User, error){
	var user *models.User
	query := bson.D{bson.E{Key:"username", Value: username}}
	err := u.usercollection.FindOne(u.ctx, query).Decode(&user)
	return user, err
}

func (u *UserServiceImpl) GetAll() ([]*models.User, error){
	var users []*models.User
	fmt.Println("here1")
	cursor, err := u.usercollection.Find(u.ctx, bson.D{{}})
	if err != nil{
		return nil, err
	}
	fmt.Println("here2")
	for cursor.Next(u.ctx){
		var user models.User
		err := cursor.Decode(&user)
		fmt.Println("here before err")
		if err != nil {
			return nil, err
		}
		fmt.Println("here after error")
		users = append(users, &user)
	}
	if err := cursor.Err(); err != nil{
		return nil, err
	}
	fmt.Println("here5")
	cursor.Close(u.ctx)

	if len(users) == 0 {
		return nil, errors.New("documents not found")
	}
	fmt.Println("here6")
	return users, nil
}

func (u *UserServiceImpl) UpdateUser(user *models.User) error{
	filter := bson.D{bson.E{Key:"user_name", Value: user.Name}}
	update := bson.D{bson.E{Key:"$set", Value: bson.D{bson.E{Key:"user_name", Value: user.Name}, bson.E{Key:"user_age", Value: user.Age}}}}
	result,_ := u.usercollection.UpdateOne(u.ctx, filter, update)
	if result.MatchedCount != 1{
		return errors.New("no match found for update")
	}
	return nil
}

func (u *UserServiceImpl) DeleteUser(name *string) error{
	// filter := bson.D{bson.E{Key:"user_name", Value: name}}
	// result, _ = u.usercollection.DeleteOne(u.ctx, filter)
	// if result.DeletedCount != 1{
	// 	return errors.New("no match found for delete")
	// } 	
	return nil
}
