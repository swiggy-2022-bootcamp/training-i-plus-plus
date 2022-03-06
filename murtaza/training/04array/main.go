package main

import (
	"01helloworld/p1"
	"fmt"
)

/**
 * Arrays are rarely used in GoLang
 * Must declare size while arrray declaration
 * Even though unused, once decalred memory remains occupied
 *
 */

func DeclareArray() {

	// Declare Array using make
	arr1 := make([]string, 10)

	arr1[0] = "apple"
	arr1[1] = "banana"
	arr1[2] = "mango"

	fmt.Println("Contents: ", arr1, "Length: ", len(arr1))

	// Declare and Initialize array
	arr2 := [5]float32{10.1, 34.6, 77.9, 80.3, 11.1}
	fmt.Println("Contents: ", arr2, "Length: ", len(arr2))
}

func main() {
	DeclareArray()
	p1.HelloWorld()
	p1.ByeWorld()
}
