package main

import "fmt"

func passByValue(a int) {
	a = 1000
}

func passByReference(ptr *int) {
	*ptr = 3000
}

func main() {
	var ptr *int
	a := 10
	ptr = &a

	passByValue(a)
	fmt.Println("Pass by value: ", a)
	passByReference(ptr)
	fmt.Println("Pass by reference: ", *ptr)
}
