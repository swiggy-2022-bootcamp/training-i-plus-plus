package main

import "fmt"

var integerSlice []int
var stringSlice []string

func main() {

	integerSlice = []int{1, 2, 3, 4, 5}
	fmt.Println(integerSlice)

	stringSlice = []string{"a", "b", "c", "d", "e"}
	fmt.Println(stringSlice)

	printInteger()
}

func printInteger() {
	fmt.Println(integerSlice)
}
