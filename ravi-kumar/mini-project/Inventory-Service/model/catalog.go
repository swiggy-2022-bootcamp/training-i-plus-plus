package mockdata

import "go.mongodb.org/mongo-driver/bson/primitive"

//import ("fmt")

type Product struct {
	Id           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name         string             `json: "name" bson: "name"`
	Price        int                `json: "price" bson: "price"`
	Description  string             `json: "description" bson: "description"`
	Seller       string             `json: "seller" bson: "seller"`
	Rating       float32            `json: "rating" bson: "rating"`
	Review       []string           `json: "review" bson: "review"`
	QuantityLeft int                `json: "quantityLeft" bson: "quantityLeft"`
}

// var Catalog = []Product{
// 	{
// 		//Id:          1,
// 		Name:        "Keyboard",
// 		Price:       700,
// 		Description: "Mechanical Keyboard",
// 		Seller:      "Zebronics",
// 		Rating:      4.2,
// 		Review:      []string{"Great product", "Worth the money"},
// 	},
// 	{
// 		//Id:          2,
// 		Name:        "Fan",
// 		Price:       1700,
// 		Description: "Cooler Fan",
// 		Seller:      "Usha",
// 		Rating:      3.9,
// 		Review:      []string{"Average rotation cycles", "Not worth the amount"},
// 	},
// 	{
// 		//Id:          3,
// 		Name:        "Shoes",
// 		Price:       3700,
// 		Description: "Running shoes",
// 		Seller:      "Puma",
// 		Rating:      4.9,
// 		Review:      []string{"Great product", "Worth the money"},
// 	},
// 	{
// 		//Id:          4,
// 		Name:        "Jeans Pant",
// 		Price:       1300,
// 		Description: "Pant for Men",
// 		Seller:      "Lee Cooper",
// 		Rating:      4.5,
// 		Review:      []string{"Shrinks quickly", "Great quality cloth"},
// 	},
// }

// //exportable function
// func GetProductCatalog() []Product {
// 	return Catalog
// }
