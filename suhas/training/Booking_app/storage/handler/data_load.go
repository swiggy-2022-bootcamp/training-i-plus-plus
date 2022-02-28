package storage

import (
	"time"
	"booking_app/admin"
	"booking_app/ticket"
	"booking_app/train"
	//"booking_app/customer"
)

var Admin_array = []admin.Admin {
	{"admin1",101},
	{"admin2",102},
	{"admin3",103},
}

var Train_array = []train.Train {
	{201,"Mysore","Banglore",time.Date(2022, 2, 1, 3, 0, 0, 0, time.UTC),time.Date(2022, 2, 1, 3, 3, 0, 0, time.UTC)},
	{202,"Banglore","Chennai",time.Date(2022, 2, 1, 3, 0, 0, 0, time.UTC),time.Date(2022, 2, 1, 3, 3, 0, 0, time.UTC)},
	{203,"Mysore","Chennai",time.Date(2022, 2, 1, 3, 0, 0, 0, time.UTC),time.Date(2022, 2, 1, 3, 3, 0, 0, time.UTC)},
}

var Ticket_array = []ticket.TicketAvailable {
	{201,100},
	{202,150},
	{203,75},
}	

func GetAdminArray () []admin.Admin {
	return Admin_array
}

func GetTrainArray () []train.Train {
	return Train_array
}

func GetTicketArray () []ticket.TicketAvailable {
	return Ticket_array
}

func AddAdminArray (Ad admin.Admin) bool {
	Admin_array = append(Admin_array,Ad)
	return true
}

func AddTrainArray (Tr train.Train) bool {
	Train_array = append(Train_array,Tr)
	return true
}

func AddTicketArray (Ti ticket.TicketAvailable) bool {
	Ticket_array = append(Ticket_array,Ti)
	return true
}