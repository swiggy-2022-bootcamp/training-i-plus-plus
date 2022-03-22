package main

import "fmt"

func main() {

	s := []int{10, 20, 30, 40, 50}

	for key, value := range s {
		fmt.Println(key, value)
	}

	for _, value := range s {
		fmt.Println(value)
	}

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

}
