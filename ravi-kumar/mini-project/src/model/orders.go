package mockdata

import "time"

type Order struct {
	UserId       string    `json: "userId" bson: "userId"`
	Items        []string  `json: "items" bson: "items"`
	Amount       float64   `json: "amount" bson: "amount"`
	OrderDate    time.Time `json: "orderDate" bson: "orderDate"`
	DeliveryDate time.Time `json: "deliveryDate" bson: "deliveryDate"`
}
