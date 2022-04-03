package main

import (
	"fmt"
	"order.akash.com/db"
	"order.akash.com/kafka"
	"order.akash.com/server"
)

func init() {
	repo := db.NewMongoRepository()
	repo.Connect()
	go kafka.StartOrderListener(repo)
	server.Start(repo)
}

func main() {
	fmt.Println("Hello from order service")
}
