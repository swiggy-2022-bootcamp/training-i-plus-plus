/*
Setting value of interface using value

*/

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 2.3
	v := reflect.ValueOf(x)
	// v.SetFloat(2.3234) // Panic: reflect.value.SetFloat using unaddressable value

	fmt.Println("Settability of v: ", v.CanSet()) // false

	// for make it settable

	// Step 1:
	v2 := reflect.ValueOf(&x) // By reference

	// Step 2:
	//create elem that points to the value of interface
	vEle := v2.Elem()

	// Step 3:
	// Check value of settablity
	fmt.Println("Settability of v2: ", vEle.CanSet()) // true

	// Step 4:
	vEle.SetFloat(2.3234)
	fmt.Println(vEle.Interface())
	fmt.Println(vEle)
}
