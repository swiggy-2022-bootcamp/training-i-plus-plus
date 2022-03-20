package ticket

import (
	scribble "github.com/nanobox-io/golang-scribble"
	"strconv"
	"fmt"
	"BookingApp/customer"
	"BookingApp/train"
	"encoding/json"
)

func IsTicketAvialable(train_id int) bool {
	for _,v := range(Ticket_array) {
		if v.Train_id == train_id && v.Count > 0 {
			v.Count -=1
			return true
		}	
	}
	return false
}


func CreateTicket(cust customer.Customer,train train.Train) TicketCustomer {
	
	var Ticket_new TicketCustomer
	Ticket_new = TicketCustomer{
		train.Train_id,
		train.Source,
		train.Destination,
		train.Arrival_time,
		train.Departure_time,
		cust.Customer_id,
	}
	return Ticket_new
}

func AddTicketBooked(Tick TicketCustomer) bool {
	var tiket_dir string = "./ticket/ticket_json"
	ticket_db, err := scribble.New(tiket_dir, nil)
	if err != nil {
	fmt.Println("Error", err)
	}
	if err:= ticket_db.Write("ticket_booked", strconv.Itoa(Tick.Customer_id)+strconv.Itoa(Tick.Train_id),Tick);err!= nil {
		//log error
		fmt.Println(err)
		return false
	}
	return true
}

func GetTicketBooked(customer_id int,train_id int) TicketCustomer {
	var tiket_dir string = "./ticket/ticket_json"
	ticket_db, err := scribble.New(tiket_dir, nil)
	if err != nil {
		fmt.Println("Error", err)
	}
	ticket_obj := TicketCustomer{}
	if err:= ticket_db.Read("ticket_booked",strconv.Itoa(customer_id)+strconv.Itoa(train_id), &ticket_obj);err!= nil {
		//log error
		fmt.Println(err)
	}
	return ticket_obj
}

func GetAllTicketsBooked() []TicketCustomer {	
	var tiket_dir string = "./ticket/ticket_json"
	ticket_db, err := scribble.New(tiket_dir, nil)
	if err != nil {
	fmt.Println("Error", err)
	}
	ticket_objs ,err := ticket_db.ReadAll("ticket_booked")
	if err!= nil {
		//log error
		fmt.Println(err)
	}

	var tickets = []TicketCustomer{}
	for _, f := range ticket_objs {
	ticketFound := TicketCustomer{}
	if err := json.Unmarshal([]byte(f), &ticketFound); err != nil {
		fmt.Println("Error", err)
	}
	tickets = append(tickets, ticketFound)
	}
	return tickets
	
}

func BookTrain(cust customer.Customer,tr train.Train) bool {
	//json
	if(IsTicketAvialable(tr.Train_id)==true){
		//add log
		return false
	}
	Tk := CreateTicket(cust,tr)
	AddTicketBooked(Tk)
	return true
}

func GetTicket(cust customer.Customer,train_id int)TicketCustomer {
	//json
	var tk TicketCustomer
	tk = GetTicketBooked(cust.Customer_id,train_id)
	return tk
}