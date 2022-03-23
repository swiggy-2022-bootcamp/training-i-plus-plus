package main

import (
	. "fmt"
	_ "net/http"
)

var num = 19

func init() {
	Println("Yellowww", num)
}

func main() {
	Print("Hello world")

	var a, b = func() (a, b int) {
		Println("Inside anonymous func")
		a = 10
		b = 12
		return
	}()

	Println("a = ", a, " b = ", b)

	num := 1
	str := "ravi"
	boolVal := true

	//%v - any variable in string form, %s - string variable in string form, %T - type of a variable, %% - for %
	Printf("num = %v, str = %s, boolVal type = %T", num, str, boolVal)
}

func init() {
	Println("Blue")
}
