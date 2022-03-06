package main

import (
	"fmt"
	"sample.akash.com/db"
	"sample.akash.com/model"
	"sample.akash.com/server"
	"sample.akash.com/user"
)

func init() {
	db.Connect()
	server.Start()
}

func main() {
	fmt.Println("Hello ", &model.User{"Ash", "Lambert", "ash.lambert@swiggy.com", "12345"})
	user.Login("abc", "def")
}
