package main

import (
	"sample.akash.com/db"
	"sample.akash.com/log"
	"sample.akash.com/model"
	"sample.akash.com/server"
)

func init() {
	db.Connect()
	server.Start()
}

func main() {
	user1 := &model.User{"Ash", "Lambert", "ash.lambert@swiggy.com", "12345"}
	log.Info("Hello ", *user1)
}
