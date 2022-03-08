package main

import (
	"fmt"

	"github.com/svett/golang-design-patterns/creational-patterns/singleton/db"
)

func main() {
	db.Repository().Set("id","ID001")
	value, err := db.Repository().Get("id")
	if err != nil {
		panic(err)
	}

	fmt.Println(value)
}