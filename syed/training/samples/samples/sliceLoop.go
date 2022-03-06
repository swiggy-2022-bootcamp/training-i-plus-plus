package main

import "fmt"

func main() {

	s := []int{10, 20, 30, 40}

	// We can loop through this slice in two ways:

	// 1. using "range"
	for key, value := range s {
		fmt.Println(key, value)
	}
	//If we dont want the key, we do, replace "key" with "_":
	for _, value := range s {
		fmt.Println(value)
	}

	// 2. Using traditional forloop:
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i]) //get the value at index "i"
	}
}