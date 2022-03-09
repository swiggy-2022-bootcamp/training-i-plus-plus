package storage

import (
	"booking_app/customer"
	"booking_app/ticket"
	"fmt"
	"strconv"
	//"encoding/json"
	scribble "github.com/nanobox-io/golang-scribble"

)

func AddCustomer(Cust customer.Customer) bool {
	var customer_dir string = "./storage/json_files/customer"
	customer_db, err := scribble.New(customer_dir, nil)
	if err != nil {
	fmt.Println("Error", err)
	}
	if err:= customer_db.Write("customer", strconv.Itoa(Cust.Customer_id), Cust);err != nil {
		//log error
		fmt.Println(err)
		return false
	}
	return true
}

func GetCustomer(id int) customer.Customer {
	var customer_dir string = "./storage/json_files/customer"
	customer_db, err := scribble.New(customer_dir, nil)
	if err != nil {
	fmt.Println("Error", err)
	}
	customer_obj := customer.Customer{}
	if err:= customer_db.Read("customer", strconv.Itoa(id), &customer_obj);err!= nil {
		//log error
		fmt.Println(err)
	}
	return customer_obj
}	

func AddTicketBooked(Tick ticket.TicketCustomer) bool {
	var tiket_dir string = "./storage/json_files/ticket_booked"
	ticket_db, err := scribble.New(tiket_dir, nil)
	if err != nil {
	fmt.Println("Error", err)
	}
	if err:= ticket_db.Write("ticket_booked", strconv.Itoa(Tick.Customer_id),Tick);err!= nil {
		//log error
		fmt.Println(err)
		return false
	}
	return true
}

func GetTicketBooked(customer_id int) ticket.TicketCustomer {
	var tiket_dir string = "./storage/json_files/ticket_booked"
	ticket_db, err := scribble.New(tiket_dir, nil)
	if err != nil {
	fmt.Println("Error", err)
	}
	ticket_obj := ticket.TicketCustomer{}
	if err:= ticket_db.Read("customer",strconv.Itoa(customer_id), &ticket_obj);err!= nil {
		//log error
		fmt.Println(err)
	}
	return ticket_obj
}