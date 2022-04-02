package model

type CartProduct struct {
	ProductId string `json:"productId" bson:"productId" example:"623f0eae80d6e879d6"`
	Quantity  int    `json:"quantity" bson:"quantity" example:"2"`
}

type AuthResponse struct {
	Text string `json:"text" example:"You are not authenticated"`
}
