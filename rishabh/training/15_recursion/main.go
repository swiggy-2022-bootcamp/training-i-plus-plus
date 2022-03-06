package main

import "fmt"

func fact(n int) int{
	if n==0{
		return 1
	}
	return n*fact(n-1)
}

func main(){
	fmt.Println("factorial of 9 is",fact(9))
}