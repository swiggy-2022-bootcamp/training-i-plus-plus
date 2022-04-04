package models
import "go.mongodb.org/mongo-driver/bson/primitive"

type Cart struct {
	CartId    primitive.ObjectID `json:"cartid,omitempty"`
	ProductId string `json:"productid" validate:"required"`
	Price     int    `json:"price" validate:"required"`
	UserId    string `json:"userid" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required"`
}
