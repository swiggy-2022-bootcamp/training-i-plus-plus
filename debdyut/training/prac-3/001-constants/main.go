package main

import "fmt"

// Constants
// const a = 42
// const b = "James"

// Un-typed constants
// const (
// 	a = 42
// 	b = "James"
// )

// Typed constants
const (
	a int    = 42
	b string = "James"
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
}
