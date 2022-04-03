package main

import (
	"sample.akash.com/api"
	"sample.akash.com/db"
	"sample.akash.com/server"
)

func init() {
	api.InitCustomerAPI(db.NewMongoRepository())
	server.Start()
}

func main() {
}
