package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	UserName      string             `json:"username"`
	EmailId       string             `json:"email_id"`
	Password      string             `json:"password"`
	Token         *string            `json:"token"`
	Refresh_token *string            `json:"refresh_token"`
}

type LoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
