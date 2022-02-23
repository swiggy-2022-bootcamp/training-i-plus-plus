package main

import "fmt"

func main() {
	if x := 42; x == 42 {
		fmt.Println("Equal")
	} else if x > 42 {
		fmt.Println("Greater")
	} else {
		fmt.Println("Lesser")
	}
	// Below line won't work as x is out of scope
	// fmt.Println(x)
}
