package controller

import (
	"fmt"
	"src/mockdata"
	"strings"
)

func GetCatalog() []mockdata.Product {
	return mockdata.GetProductCatalog()
}

func PrintCatalog() {
	catalog := mockdata.GetProductCatalog()
	for _, product := range catalog {
		PrintProduct(&product)
	}
}

func PrintProduct(product *mockdata.Product) {
	fmt.Println("\nname: ", product.Name,
		"\nprice: ", product.Price,
		"\ndescription: ", product.Description,
		"\nseller: ", product.Seller,
		"\nrating: ", product.Rating,
		"\nreview: ", strings.Join(product.Review, ", "))
}
