package model

type Product struct {
	Id            string `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Price         string `json:"price"`
	Seller        string `json:"seller"`
	SellerContact string `json:"seller-contact"`
}
