package main

import "fmt"

func addUser(name string, age int) {
	fmt.Printf("adding user = %s, age = %d \n", name, age)
}

func main() {
	addUser("Nishant", 4)
	addUser("Neo", 12)
}
