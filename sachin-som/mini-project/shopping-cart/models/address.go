package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Address struct {
	AddressID primitive.ObjectID `json:"_id"      bson:"_id"`
	House     *string            `json:"house"    bson:"house"     validate:"required"`
	Street    *string            `json:"street"   bson:"street"    validate:"required"`
	City      *string            `json:"city"     bson:"city"      validate:"required"`
	Pincode   *string            `json:"pincode"  bson:"pincode"   validate:"required"`
	State     *string            `json:"state"    bson:"state"     validate:"required"`
	Country   *string            `json:"country"  bson:"country"   validate:"required"`
	Landmark  *string            `json:"landmark" bson:"landmark"`
}
