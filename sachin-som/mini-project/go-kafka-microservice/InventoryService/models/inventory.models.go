package models

type Inventory struct {
	ID                int       `json:"_id"                bson:"_id"                 validate:"required"`
	InventoryName     *string   `json:"inventory_name"     bson:"inventory_name"      validate:"required, min=4, max=50"`
	OwnerID           *string   `json:"owner_id"           bson:"owner_id"            validate:"required"`
	InventoryProducts []Product `json:"inventory_products" bson:"inventory_products"`
}
