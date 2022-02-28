package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	// Arrays

	var fruitList [4]string
	fruitList[0] = "apple"
	fruitList[1] = "banana"
	fruitList[3] = "dragon-fruit"

	fmt.Println(fruitList)
	fmt.Printf("Type: %T\n", fruitList)
	fmt.Println("Length: ", len(fruitList))

	var fruitList2 = [4]string{"a", "b", "c"}
	fmt.Println(fruitList2)
	fmt.Printf("Type: %T\n", fruitList2)
	fmt.Println("Length: ", len(fruitList2))

	fruitList3 := [4]string{"a", "b", "c"}
	fmt.Println(fruitList3)
	fmt.Printf("Type: %T\n", fruitList3)
	fmt.Println("Length: ", len(fruitList3))

	// Slices

	var fruitList4 = []string{}
	printInfo(fruitList4)

	fruitList5 := append(fruitList4, "ace", "king", "queen")
	printInfo(fruitList5)

	prices := make([]int, 0)
	prices = append(prices, 3, 7, 45, 22, 9, 40)
	fmt.Println(prices)

	sort.Ints(prices)
	fmt.Println(prices, sort.IntsAreSorted(prices))

}

func printInfo(list []string) {
	fmt.Println(list)
	fmt.Printf("Type: %T\n", list)
	fmt.Println("Length: ", len(list))
	fmt.Println()
}
