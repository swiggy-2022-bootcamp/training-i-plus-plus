package main 

import (
	"fmt"
	"booking_app/customer"
	"booking_app/storage/handler"
)

func main() {
	
	// create Customer
	// Check available Train
	// book train -> get ticket -> check json file
	// add the above log 
	fmt.Println("Hello")
	storage.AddLog("Hello")
	cust1 := customer.Customer{"Suhas","R",501}
	cust2 := customer.Customer{"Kiran","K",502}
	storage.AddCustomer(cust1)
	storage.AddCustomer(cust2)
}