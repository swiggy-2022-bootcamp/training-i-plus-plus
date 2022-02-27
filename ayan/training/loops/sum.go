package main

import "fmt"

func main() {
	fmt.Println("Sum is : ", add(10, 20))
}

func add(i int, j int) int {
	return i + j
}
