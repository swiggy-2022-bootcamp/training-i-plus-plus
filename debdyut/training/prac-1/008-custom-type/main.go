package main

import "fmt"

type hotdog int

var a hotdog = 5

func main() {
	// Print the hotdog custom type value
	res := fmt.Sprintf("a is of type: %T, with value: %d", a, a)
	fmt.Println(res)

	var b int
	// invalid assignment, cannot assign main.hotdog to int type
	// b = a

	// Type conversion example
	b = int(a)

	fmt.Println("b:", b)
}
