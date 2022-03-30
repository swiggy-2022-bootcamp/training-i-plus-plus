package model

type OrderStatus string

const (
	ORDER_PLACED     OrderStatus = "ORDER_PLACED"
	DISPATCHED       OrderStatus = "DISPATCHED"
	OUT_FOR_DELIVERY OrderStatus = "OUT_FOR_DELIVERY"
	CANCELLED        OrderStatus = "CANCELLED"
	ORDER_COMPLETED  OrderStatus = "ORDER_COMPLETED"
)

type AuthResponse struct {
	Text string `json:"text"`
}

type Product struct {
	ProductId string  `json:"productId" bson:"productId"`
	Quantity  int     `json:"quantity" bson:"quantity"`
	Price     float32 `json:"price" bson:"price"`
}

type Order struct {
	Email       string      `json:"email" bson:"email"`
	Products    []Product   `json:"products" bson:"products"`
	OrderStatus OrderStatus `json:"orderStatus" bson:"orderStatus"`
}

// doesn't contain user email
type UserOrder struct {
	Products    []Product   `json:"products" bson:"products"`
	OrderStatus OrderStatus `json:"orderStatus" bson:"orderStatus"`
}

type AllOrders struct {
	Orders []UserOrder `json:"orders" bson:"orders"`
}
