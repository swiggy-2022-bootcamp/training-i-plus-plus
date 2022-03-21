package models

import ()

type Appointment struct {
	Slot			string	`json:slot`
	Fees			int		`json:fees`
	Occupied		bool	`json:occupied`
}