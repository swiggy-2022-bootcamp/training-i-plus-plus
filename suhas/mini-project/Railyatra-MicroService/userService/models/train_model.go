package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Train struct {
	Id       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Station1 string             `json:"station1,omitempty" validate:"required"`
	Station2 string             `json:"station2,omitempty" validate:"required"`
}

type CheckAvialTrain struct {
	Trainid  primitive.ObjectID `json:"id" bson:"_id,omitempty" validate:"required"`
	Capacity int32              `json:"capacity" validate:"required"`
}
