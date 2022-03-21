package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Patient struct {
	Id 					primitive.ObjectID		`bson:"_id,omitempty"`
	User										`json:"user" validate:"required"`
	DoctorAssignedId	primitive.ObjectID		`json:"doctorAssignedId" validate:"required"`
	IsDischarged		bool					`json:"isDischarged" validate:"required"`
	RoomAllocated		string					`json:"roomAllocated" validate:"required"`
}