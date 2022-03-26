package main

import "fmt"

func main() {
	fmt.Println("Sum is : ", add(5, 6))
}

func add(a int, b int) int {
	return a + b
}
