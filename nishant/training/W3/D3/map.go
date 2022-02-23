package main

import "fmt"

func main() {

	myMap := make(map[string]int)

	myMap["apple"] = 1
	myMap["Orange"] = 5

	for k, v := range myMap {
		fmt.Println(k, v)
	}

	delete(myMap, "apple")
	fmt.Println("--- deleted ---")
	for k, v := range myMap {
		fmt.Println(k, v)
	}

	fmt.Println("--- 2 d ---")

	twoDMap := map[int]map[string]int{
		34: {"str": 5},
		65: {"test": 22},
	}

	for k, v := range twoDMap {
		fmt.Println(k, v)
	}

}
