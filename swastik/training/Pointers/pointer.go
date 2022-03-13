package main

import (
    "fmt"
)

func main() {
	var x int = 123

	var p * int
	p = &x

	fmt.Println("address: ",p)
	fmt.Println("value: ",*p)

	var s *int 
	fmt.Println("default value: ", s)
}
