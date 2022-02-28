package main

import "fmt"

func main(){

	var isLoggedIn bool = true

	switch isLoggedIn{
	case true:
		fmt.Println("User is logged in")
	case false:
		fmt.Println("User is logged out")
	}

	orders := 1

	switch orders{
	case 1:
		fmt.Println("One order is placed")
	case 2: 
		fmt.Println("Two orders are placed")
	case 3:
		fmt.Println("Multiple Orders can't count")
	}

}