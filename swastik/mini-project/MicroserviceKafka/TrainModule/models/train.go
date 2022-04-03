package models

type Train struct{
	Train_number 	int		`json:"train_number"`
	Train_name 		string	`json:"train_name"`
	Source			string	`json:"source"`
	Destination		string	`json:"destination"`
	Seats_available	[]int	`json:"seats_available"`
	Total_seats		int 	`json:"total_seats"`
}

type Ticket struct{
	PNR_number 		int		`json:"pnr_number"`
	Train_number 	int		`json:"train_number"`
	Seat_number		int 	`json:"seat_number"`
	Date_time		string	`json:"date_time"`
	Passenger_name	string	`json:"passenger_name"`
	Source			string	`json:"source"`
	Destination		string	`json:"destination"`
}	