package main

import (
	"fmt"
)
func main() {
	fmt.Println("Arrays in go")

	var arr [5]int
	arr[0]=10
	arr[1]=20
	arr[2]=30
	arr[3]=40
	arr[4]=50
	fmt.Println(arr)

	arr1 := [3]int{10,20,30}
	fmt.Println(arr1)
	
	fmt.Println(len(arr1))
	
    sum:=0
	for i:=0;i<len(arr);i++{
     sum+=arr[i]
	}
	fmt.Println(sum)

}