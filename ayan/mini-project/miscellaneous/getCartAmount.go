package main

import (
	"fmt"
)

type cartItem struct {
	name     string
	quantity int
}

type cart struct {
	cartList []cartItem
}

func (c cart) getAmount(priceMap map[string]int) int {

	amount := 0

	for _, item := range c.cartList {
		amount += item.quantity * priceMap[item.name]
	}
	return amount
}

func main() {

	var amount int

	priceMap := map[string]int{
		"butter":   100,
		"biscuits": 50,
		"jam":      200,
	}

	c := cart{[]cartItem{
		{"butter", 2},
		{"biscuits", 3},
		{"jam", 1}},
	}

	amount = c.getAmount(priceMap)

	fmt.Println("Total amount is ", amount)

}
