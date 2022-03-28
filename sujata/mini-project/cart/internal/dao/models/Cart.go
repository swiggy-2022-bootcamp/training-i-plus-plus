package model

type Cart struct {
	Products []CartProduct `json:"products" bson:"products"`
}
