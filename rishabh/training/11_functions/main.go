package main

import "fmt"

func add( a,b int) int{
	return a+b
}

func double( a int) int{
	return add(a,a)
}

func main(){
	a := 2
	b := 10

	fmt.Println(" a+b = ",add(a,b))
	fmt.Println(" 2*b = ",double(a))
}