package mockdata

//import ("fmt")

type Product struct {
	Name        string
	Price       int
	Description string
	Seller      string
	Rating      float32
	Review      []string
}

//exportable function
func GetProductCatalog() []Product {
	var catalog = []Product{
		{
			Name:        "Keyboard",
			Price:       700,
			Description: "Mechanical Keyboard",
			Seller:      "Zebronics",
			Rating:      4.2,
			Review:      []string{"Great product", "Worth the money"},
		},
		{
			Name:        "Fan",
			Price:       1700,
			Description: "Cooler Fan",
			Seller:      "Usha",
			Rating:      3.9,
			Review:      []string{"Average rotation cycles", "Not worth the amount"},
		},
		{
			Name:        "Shoes",
			Price:       3700,
			Description: "Running shoes",
			Seller:      "Puma",
			Rating:      4.9,
			Review:      []string{"Great product", "Worth the money"},
		},
		{
			Name:        "Jeans Pant",
			Price:       1300,
			Description: "Pant for Men",
			Seller:      "Lee Cooper",
			Rating:      4.5,
			Review:      []string{"Shrinks quickly", "Great quality cloth"},
		},
	}
	return catalog
}
