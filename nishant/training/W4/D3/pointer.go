package main

import (
	"fmt"
)

func main() {

	a := "safssf"
	var b *string = &a

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b)

	*b = "xyz"
	fmt.Println(b)
	fmt.Println(*b)

	fmt.Println("--------------")

	fmt.Println(a)

	passByValue(a)
	fmt.Println(a)

	passByRef(&a)
	fmt.Println(a)

}

func passByValue(a string) {
	a = "qwer"
	fmt.Println(a)
}

func passByRef(a *string) {
	*a = "qwer"
	fmt.Println(a)
}
