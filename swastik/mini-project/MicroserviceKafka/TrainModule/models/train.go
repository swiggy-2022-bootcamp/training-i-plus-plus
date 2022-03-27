package models

type Train struct{
	Train_number 	string	`json:"train_number"`
	Train_name 		string	`json:"train_name"`
	Source			int		`json:"source"`
	Destination		string	`json:"destination"`
	Seat_available	int		`json:"seat_available"`
	Total_seats		int 	`json:"total_seats"`
}
