package main

import (
	"fmt"
)

func main() {

	c1 := complex(1, 2)

	fmt.Println(" Value of c1 is:", c1)

	fmt.Println("Real:", real(c1))
	fmt.Println("Imaginary:", imag(c1))

}
