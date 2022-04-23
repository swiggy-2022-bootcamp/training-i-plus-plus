package main

import "fmt"

func getBackNames() (string, string) {
	return "Ashwin", "Gopalsamy"
}

func main() {
	firstName, lastName := getBackNames()
	fmt.Println("Hey there ! :)", firstName, lastName)
}
