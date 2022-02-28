package main

//Using Scribble A tiny JSON database in Golang

import (
	"BookingApp/admin"
	"BookingApp/customer"
	"BookingApp/logger"
	"BookingApp/ticket"
	"BookingApp/train"
	"fmt"
	"time"
)

func test_channel_func(cus customer.Customer, cha chan string) {
	customer.AddCustomer(cus)
	go logger.AddLog("Added a new customer")
	fmt.Println(customer.GetCustomer(301))
	fmt.Println(customer.GetCustomers())
	time.Sleep(2 * time.Second)
	cha <- "Added a new customer"
}

func main() {
	fmt.Println("Hello")
	logger.AddLog("Tested OK")

	//admin isolated test done
	newadmin := admin.Admin{
		Name:   "admin5",
		Userid: 105,
	}
	go newadmin.AddAdmin()
	go logger.AddLog("Added a new Admin")

	fmt.Println(newadmin.CheckAdmin())
	fmt.Println(admin.GetAdminArray())

	//customer isolated test done
	newcustomer := customer.Customer{
		Firstname:   "joe",
		Lastname:    "doe",
		Customer_id: 302,
	}

	sample_channel := make(chan string, 1)
	test_channel_func(newcustomer, sample_channel)

	select {
	case msg := <-sample_channel:
		fmt.Println("Recived message", msg)
	default:
		fmt.Println("No message recieved")
	}

	//ticket isolated test done
	newticket := ticket.TicketCustomer{
		Train_id:       203,
		Source:         "Mysore",
		Destination:    "Chennai",
		Arrival_time:   time.Date(2022, time.Month(2), 16, 1, 10, 30, 0, time.UTC),
		Departure_time: time.Date(2021, time.Month(2), 23, 1, 10, 30, 0, time.UTC),
		Customer_id:    302,
	}
	ticket.AddTicketBooked(newticket)
	go logger.AddLog("Added a new ticket")
	fmt.Println(ticket.GetTicketBooked(302,203))
	fmt.Println(ticket.GetAllTicketsBooked())

	//train isolated test done
	newtrain := train.Train{
		Train_id:       205,
		Source:         "Delhi",
		Destination:    "Noida",
		Departure_time: time.Date(2022, 2, 1, 3, 0, 0, 0, time.UTC),
		Arrival_time:   time.Date(2022, 2, 1, 3, 3, 0, 0, time.UTC),
	}

	train.AddTrain(newtrain)
	fmt.Println(train.GetTrainArray())

	//add ticket
	newticketavail := ticket.TicketAvailable{
		Train_id: 203,
		Count:    50,
	}
	newadmin.AddTicketAdmin(newticketavail)

	//get train
	fmt.Println(train.GetTrainArray())
	//book ticket
	ticket.BookTrain(newcustomer, newtrain)
	fmt.Println(ticket.GetTicket(newcustomer,newticketavail.Train_id))
}
