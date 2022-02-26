package main

import "fmt"

type A struct{ a, b int }
type B struct{ a, b, c int }
type c struct {
	A
	B
}

func main() {
	cs := c{A{a: 3, b: 4}, B{a: 1, b: 8, c: 0}}
	// fmt.Println(cs.a)  // ambiguous selector cs.acompilerAmbiguousSelector
	fmt.Println(cs)
}
