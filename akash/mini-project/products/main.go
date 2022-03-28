package main

import (
	"products.akash.com/db"
	"products.akash.com/log"
	"products.akash.com/model"
	"products.akash.com/server"
)

func init() {
	db.Connect()
	server.Start()
}

func main() {
	product1 := &model.Product{"WL1006", "Carbon Fibre Knife", "A compact knife made out of carbon fibre", "1000", "Lexcorp", "932312345"}
	log.Info("Hello ", *product1)
}
