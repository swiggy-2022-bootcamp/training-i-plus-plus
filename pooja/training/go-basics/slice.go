package main

import "fmt"

var integerSlice []int
var stringSlice []string

func main() {
	integerSlice = []int{10, 20, 30, 40}
	fmt.Println("this is integer slice ", integerSlice)

	stringSlice = []string{"first", "second", "third"}
	fmt.Println("this is string slice ", stringSlice)

	printInteger()
}

func printInteger() {
	fmt.Println("Print the numbers ", integerSlice)
}
