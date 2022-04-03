package model

type Order struct {
	Username  string `json:"username"`
	ProductId string `json:"product-id"`
	Quantity  string `json:"quantity"`
	Time      string `json:"time"`
}
