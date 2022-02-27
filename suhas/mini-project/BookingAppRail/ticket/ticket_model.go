package ticket

import (
	"time"
	//"BookingApp/customer"
	//"BookingApp/train"
)

type TicketAvailable struct {
	Train_id  int 
	Count     int
}

type TicketCustomer struct {
	Train_id  		int
	Source   		string
	Destination     string
	Arrival_time    time.Time
	Departure_time  time.Time
	Customer_id     int
}
var Ticket_array = []TicketAvailable {
	{
		Train_id:201,
		Count:100,
	},
	{
		Train_id:202,
		Count:150,
	},
	{
		Train_id:203,
		Count:75,
	},
}	

