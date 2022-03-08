package main

import "fmt"

var integerSlice []int
var stringSlice []string

func main(){
	integerSlice := []int{10,20,30,40}
	stringSlice := []string{"ten","twenty","thirty","forty"}
	fmt.Println(integerSlice, stringSlice)
}