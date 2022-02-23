package main

import "fmt"

// package level scope
var x int
var y string
var z bool

func main() {
	// Print default values
	fmt.Printf("%T has a default value of %d\n", x, x)
	fmt.Printf("%T has a default value of %s\n", y, y)
	fmt.Printf("%T has a default value of %t\n", z, z)
}
