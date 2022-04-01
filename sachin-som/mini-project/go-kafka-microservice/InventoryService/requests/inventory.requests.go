package requests

type InventoryRequest struct {
	InventoryName *string `json:"inventory_name"     bson:"inventory_name"      validate:"required, min=4, max=50"`
	OwnerID       *string `json:"owner_id"           bson:"owner_id"            validate:"required"`
}

type ProductRequest struct {
	ProductName string `json:"product_name" bson:"product_name"`
	Description string `json:"description"  bson:"description"`
	Quantity    string `json:"quantity"     bson:"quantity"`
	Price       string `json:"price"        bson:"price"`
	Ratings     uint   `json:"ratings"      bson:"ratings"`
	ImageUrl    string `json:"image_url"    bson:"image_url"`
}
