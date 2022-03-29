package model

type Product struct {
	Name        string  `json:"name" bson:"name"`
	Description string  `json:"description" bson:"description"`
	Price       float32 `json:"price" bson:"price"`
	Quantity    int     `json:"quantity" bson:"quantity"`
}

type ProductIds struct {
	Ids []string `json:"ids"`
}

type Products struct {
	Products []Product `json:"products"`
}
