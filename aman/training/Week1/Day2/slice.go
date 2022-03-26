package main

import "fmt"

var integerSlice []int
var stringSlice []string

func main() {
	integerSlice = []int{1, 2, 3, 4}
	fmt.Println(integerSlice)

	stringSlice = []string{"first", "second"}
	fmt.Println(stringSlice)

	printInteger()
}

func printInteger() {
	fmt.Println("Print the numbers ", integerSlice)
}
