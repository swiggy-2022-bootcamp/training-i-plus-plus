package main

import "fmt"

func main() {

	arr := [5]int{1, 2, 3, 4, 5}

	for i, val := range arr {
		fmt.Println(val, " at ", i)
	}

}
