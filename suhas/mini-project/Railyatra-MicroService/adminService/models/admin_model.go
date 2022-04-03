package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// type Admin struct {
//     Id       primitive.ObjectID `json:"id,omitempty"`
//     Name     string             `json:"name,omitempty" validate:"required"`
//     Location string             `json:"location,omitempty" validate:"required"`
//     Title    string             `json:"title,omitempty" validate:"required"`
// }

type Admin struct {
	Id    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name  string             `json:"name,omitempty" validate:"required"`
	Email string             `json:"email,omitempty" validate:"required"`
}
