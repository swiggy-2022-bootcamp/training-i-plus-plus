package main

import (
	"fmt"
)

type product interface {
	getDiscountedPrice() float64
	checkAvailability(int) bool
	// updateAvailability(int)
}

type food struct {
	name              string
	price             float64
	availableQuantity int
	discount          float64
}

func (f food) getDiscountedPrice() float64 {
	return f.price * (1 - f.discount/100)
}

func (f food) checkAvailability(requiredQuantity int) bool {
	return f.availableQuantity >= requiredQuantity
}

// func (f food) updateAvailability(requiredQuantity int) {
// 	f.availableQuantity -= requiredQuantity
// }

type cartItem struct {
	name     string
	quantity int
}

type cart struct {
	cartList []cartItem
}

func (c cart) checkout(inventoryMap map[string]product) float64 {

	amount := 0.0

	for _, item := range c.cartList {
		pdt := inventoryMap[item.name].(food)

		if pdt.checkAvailability(item.quantity) {

			amount += float64(item.quantity) * pdt.getDiscountedPrice()
			// pdt.updateAvailability(item.quantity)
		}

	}
	fmt.Println(inventoryMap)

	return amount
}

func main() {

	var amount float64

	inventoryMap := map[string]product{
		"butter": food{
			name:              "butter",
			price:             100,
			availableQuantity: 10,
			discount:          10,
		},
		"biscuits": food{
			name:              "biscuits",
			price:             100,
			availableQuantity: 10,
			discount:          50,
		},
		"jam": food{
			name:              "jam",
			price:             200,
			availableQuantity: 10,
			discount:          20,
		},
	}

	c := cart{[]cartItem{
		{"butter", 2},
		{"biscuits", 3},
		{"jam", 1}},
	}

	// fmt.Println(inventoryMap)

	amount = c.checkout(inventoryMap)

	// fmt.Println(inventoryMap)

	fmt.Println("Total amount is ", amount)

}
