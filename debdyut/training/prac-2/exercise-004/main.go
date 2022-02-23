package main

import "fmt"

// mytype has an underlying type of int
type mytype int

var x mytype

func main() {
	fmt.Printf("x is of type: %T, and has a value of: %v\n", x, x)

	x = 42

	fmt.Printf("x is of type: %T, and has a value of: %v\n", x, x)
}
