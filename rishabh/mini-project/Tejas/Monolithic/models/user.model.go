package models

import (
	"tejas/configs"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id"`
	Name     string             `json:"name" bson:"name"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
	IsAdmin  bool               `json:"is_admin" bson:"is_admin"`
}

var UserCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
