package main

import "fmt"

func zeroByValue(i int){
	i = 0
}

func zeroByReference(i *int){
	*i = 0
}

func main(){
	i:= 1
	j:= 1

	zeroByValue(i)
	zeroByReference(&i)
	fmt.Println(i,j)
}