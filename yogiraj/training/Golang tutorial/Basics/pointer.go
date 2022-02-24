package main

import "fmt"

func main(){
	var creature string = "shark"
	var pointer *string = &creature
	fmt.Println(pointer)
	fmt.Println(*pointer)
}