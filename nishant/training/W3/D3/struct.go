package main

import "fmt"

type User struct {
	name string
	age  int
}

func main() {
	musr := User{"abc", 123}
	fmt.Println(musr)
	fmt.Println(musr.name)

}
