package main

import "fmt"

// Package level scope
// DECLARE the variable "y"
// ASSIGN the value 43
// declare & assign = initialization
var y = 43

// DECLARE there is a VARIABLE with the identifier "z"
// and that the VARIABLE with the IDENTIFIER "z" is of TYPE int
// ASSIGNS the ZERO value of TYPE int to "z"
// false for booleans, 0 for integers, 0.0 for floats, "" for strings
// and nil for pointers, functions, interfaces, slices, channels, and maps.
var z int

func main() {
	// short declaration operator
	// DECLARE a variable and ASSIGN a VALUE (of a certain type)
	x := 42
	fmt.Println("x:", x)
	foo()

	fmt.Println("z:", z)
}

func foo() {
	fmt.Println("y:", y)
}
