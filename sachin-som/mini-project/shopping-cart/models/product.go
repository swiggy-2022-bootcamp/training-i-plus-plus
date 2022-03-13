package models

import (
	"time"
)

type Product struct {
	ProductID          int       `json:"_id"          bson:"_id"`
	ProductName        *string   `json:"product_name" bson:"product_name"`
	ProductDescription *string   `json:"description"  bson:"description"`
	Price              int       `json:"price"        bson:"price"`
	Ratings            *uint     `json:"ratings"      bson:"ratings"`
	ImageUrl           *string   `json:"image_url"    bson:"image_url"`
	UserId             int       `json:"upload_by"    bson:"upload_by"`
	InventoryId        int       `json:"inventory_id" bson:"inventory_id"`
	UploadedAt         time.Time `json:"uploaded_at"  bson:"uploaded_at"`
	UpdatededAt        time.Time `json:"updated_at"   bson:"updated_at"`
}
