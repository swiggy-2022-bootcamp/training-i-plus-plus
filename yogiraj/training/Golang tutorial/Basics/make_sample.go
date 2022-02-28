package main

import "fmt"

func main(){
	var nums = make([]int, 4,7)
	fmt.Printf("numbers = %v \nlength=%d \ncapacity=%d", nums,len(nums),cap(nums))
	
}