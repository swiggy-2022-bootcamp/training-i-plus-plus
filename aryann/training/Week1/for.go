package main

import "fmt"

// := syntax is shorthand for declaring and initializing a variable.

func main() {

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println("Sum:", sum)
}
