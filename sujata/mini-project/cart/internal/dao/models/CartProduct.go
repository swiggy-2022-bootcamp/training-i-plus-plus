package model

type CartProduct struct {
	ProductId string  `json:"productId" bson:"productId"`
	Quantity  int     `json:"quantity" bson:"quantity"`
	Price     float32 `json:"price" bson:"price"`
}

type AuthResponse struct {
	Text string `json:"text"`
}
