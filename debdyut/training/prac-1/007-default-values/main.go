package main

import "fmt"

var a int
var b float64
var c string
var d bool

func main() {
	// DEFAULT values for some PRIMITIVE types
	fmt.Printf("a is of type: %T, with default value: %d\n", a, a) // 0
	fmt.Printf("b is of type: %T, with default value: %f\n", b, b) // 0.0
	fmt.Printf("c is of type: %T, with default value: %s\n", c, c) // ""
	fmt.Printf("d is of type: %T, with default value: %t\n", d, d) // false
}
