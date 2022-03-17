package models

type GeneralUser struct{
	Id 					string		`json:"id"`
	User							`json:"user"`
	PreviousDiseases	string		`json:"previousDisease"`
	IsPatient			bool		`json:"isPatient"`	

}

