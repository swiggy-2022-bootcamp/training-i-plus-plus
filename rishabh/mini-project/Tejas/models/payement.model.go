package models

type Payement struct {
	Transaction_id string `json:"transaction_id" bson:"transaction_id"`
	User_id        string `json:"user_id" bson:"user_id"`
	Amount         int    `json:"amount" bson:"amount"`
	Status         string `json:"status" bson:"status"`
}
