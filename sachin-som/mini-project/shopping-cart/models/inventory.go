package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Inventory struct {
	InventoryID       primitive.ObjectID `json:"_id"                bson:"_id"                 validate:"required"`
	InventoryName     *string            `json:"inventory_name"     bson:"inventory_name"      validate:"required, min=4, max=50"`
	InventoryOwner    *string            `json:"inventory_owner"    bson:"inventory_owner"     validate:"required"`
	InventoryProducts []Product          `json:"inventory_products" bson:"inventory_products"`
	CreatedAt         time.Time          `json:"created_at"         bson:"created_at"`
	UpdatedAt         time.Time          `json:"updated_at"         bson:"updated_at"`
}
