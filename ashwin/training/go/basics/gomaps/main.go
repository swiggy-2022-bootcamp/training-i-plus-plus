package main

import "fmt"

func main() {
	products := make(map[string]string)

	products["catergory-1"] = "quantity-1"
	products["catergory-2"] = "quantity-2"
	products["catergory-3"] = "quantity-3"

	fmt.Println("products :", products)
	catergory4, isPresent := products["catergory-4"]
	fmt.Println(isPresent, catergory4)
}
