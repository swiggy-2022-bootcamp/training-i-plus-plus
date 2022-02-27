package ticket 

import "time"

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

// func AddTickets(tk TicketCustomer) bool {

// }

// func GetTickets() TicketCustomer[] {
	
// }