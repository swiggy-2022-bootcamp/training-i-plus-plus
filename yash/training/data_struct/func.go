package main

import "fmt"

func authenticate_admin(username string, password string) bool {
	if username == "admin" && password == "admin" {
		return true
	}
	return false
}

func main() {

	var integerslice []int
	integerslice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Print("The integer slice is", integerslice)
	fmt.Print("\n")
	var stringslice []string
	stringslice = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	fmt.Print("The string slice is", stringslice)
	fmt.Print("\n")

}
