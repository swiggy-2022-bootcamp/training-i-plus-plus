package models
import "go.mongodb.org/mongo-driver/bson/primitive"


type Order struct{
	OrderId   primitive.ObjectID `json:"orderid,omitempty"`
	ProductId string `json:"productid" validate:"required"`
	SellerId  string `json:"sellerid" validate:"required"`
	Status    string `json:"status" validate:"required"`
	Price     int    `json:"price" validate:"required"`
	UserId    string `json:"userid" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required"`
	Total     int    `json:"total" validate:"required"`
}