package main

import "fmt"

func main() {

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("Go Loops")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	var isLoggedIn bool = true
	if isLoggedIn {
		fmt.Println("Logged In")
	} else {
		fmt.Println("Not Logged In")
	}

	shoppingQueue := 1

	switch shoppingQueue {
	case 1:
		fmt.Println("Standing first in line")
	case 2:
		fmt.Println("Standing second in line")
	case 3:
		fmt.Println("Standing third in line")
	}

	arrRange := [5]int{1, 2, 3, 4, 5}

	for i, item := range arrRange {
		fmt.Println(item, " at ", i)
	}
}
