package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Inventory struct {
	ID            primitive.ObjectID   `json:"_id"                bson:"_id"                 validate:"required"`
	InventoryName *string              `json:"inventory_name"     bson:"inventory_name"      validate:"required, min=4, max=50"`
	OwnerID       *string              `json:"owner_id"           bson:"owner_id"            validate:"required"`
	Products      []primitive.ObjectID `json:"products" bson:"products"`
}
