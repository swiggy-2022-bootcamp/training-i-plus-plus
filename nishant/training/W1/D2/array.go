package main


import "fmt"


var arr1 [4] int;


func main() {

	arr1[0] = 2
	arr1[2] = 60

	fmt.Println("arr1 ", arr1)

	arr2 := [3]string{ "a1", "a2", "a3"}
	fmt.Println("arr2 ", arr2)


	twoDarr := [2][2]int {{3, 4}, {7, 5}}
	fmt.Println("2 d arr ", twoDarr)

	arr3 := arr2[1:3]
	fmt.Println("arr3 ", arr3)

	stringSlice := []string { "s1", "s2", "s3", "s4", "s5"}
	fmt.Println("str slice ", stringSlice, cap(stringSlice))

	fmt.Println("slice append ", append(stringSlice, "s10", "s20"))
}

