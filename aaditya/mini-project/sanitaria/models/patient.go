package models

type Patient struct {
	Id 					string		`json:"id"`
	User							`json:"user"`
	DoctorAssignedId	string		`json:"doctorAssignedId"`
	IsDischarged		bool		`json:"isDischarged"`
	RoomAllocated		string		`json:"roomAllocated"`
}