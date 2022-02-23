package main

import "fmt"

func main() {
	// alias for uint8
	var a byte = 42

	// alias for int32
	var b rune = 236789

	fmt.Printf("a is of type: %T, and has value: %d\n", a, a)
	fmt.Printf("b is of type: %T, and has value: %d\n", b, b)
}
