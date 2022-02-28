package main

import "fmt"

func main(){
	x := 1
	y := &x
	fmt.Println(*y)
	*y = 2
	fmt.Println(x)
}