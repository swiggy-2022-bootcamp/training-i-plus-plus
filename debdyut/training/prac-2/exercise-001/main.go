package main

import "fmt"

func main() {
	// ASSIGN values using SHORT DECLARATION
	x := 42
	y := "James Bond"
	z := true

	// Print values in multiple statements
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

	// Print values single statement
	res := fmt.Sprintf("x: %d; y: %s; z: %t", x, y, z)
	fmt.Println(res)

}
