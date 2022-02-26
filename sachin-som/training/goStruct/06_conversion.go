/**
When we have a struct type and
define an alias type for it
both types have the same
underlying type and can be converted into
one another
*/

package main

import "fmt"

// My Number struct
type myfloat32_1 struct {
	val float32
}

// Alias type
type myfloat32_2 myfloat32_1

func main() {
	a := myfloat32_1{5.0}
	b := myfloat32_2{2.0}

	var c = myfloat32_1(b)

	fmt.Println(a, c)
	fmt.Printf("%T", c)

}
