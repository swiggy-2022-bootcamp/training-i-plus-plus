package main

import (
	"fmt"
)



func main() {
	
	menu := map[string]float64{
		"soup":4.5,
		"salad":10,
	}

	fmt.Print(menu["soup"])

}