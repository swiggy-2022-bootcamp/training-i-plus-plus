package main

import "fmt"

func main() {
	var a int
	fmt.Println("The default value of a is:", a)

	var b float64
	fmt.Println("The default value of b is:", b)

	var c string
	fmt.Println("The default value of c is:", c)

	var d complex128
	fmt.Println("The default value of d is:", imag(d))

	var e chan interface{}
	fmt.Println("The default value of e is:", e)

	s := 'a'
	fmt.Println("The default value of s is:", rune(s))
}
