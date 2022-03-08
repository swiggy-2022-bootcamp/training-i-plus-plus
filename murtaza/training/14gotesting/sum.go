package main

import "fmt"

func Add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("Sum: ", Add(1, 2))
}
