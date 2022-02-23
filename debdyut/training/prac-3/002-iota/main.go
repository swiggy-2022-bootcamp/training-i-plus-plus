package main

import "fmt"

// Automatic increment
const (
	a = iota
	b
	c
)

const (
	d = iota
	e = iota + 1
	f
)

func main() {
	fmt.Printf("%T\n", a)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
