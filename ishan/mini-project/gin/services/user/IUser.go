package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID             primitive.ObjectID `bson:"_id",omitempty`
	Username       string             `bson:"username"`
	HashedPassword string             `bson:"password"`
	Role           string             `bson:"role"`
}

type UserPublic struct {
	ID       primitive.ObjectID `bson:"_id",omitempty`
	Username string             `bson:"username"`
}
