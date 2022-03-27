package models

type Ticket struct{
	PNR_number 		string	`json:"pnr_number"`
	Train_number 	string	`json:"train_number"`
	Seat_number		string 	`json:"seat_number"`
	Date_time		string	`json:"date_time"`
	Passenger_name	string	`json:"passenger_name"`
	Source			string	`json:"source"`
	Destination		string	`json:"destination"`

}		
