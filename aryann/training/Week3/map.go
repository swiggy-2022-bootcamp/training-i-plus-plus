package main

import "fmt"

var nameAgeMap map[string]int

func main() {

	nameAgeMap = map[string]int{
		"John": 25,
		"Mike": 30,
	}

	fmt.Println(nameAgeMap["John"])

	delete(nameAgeMap, "John")

	for k, v := range nameAgeMap {
		fmt.Println(k, v)
	}

}
