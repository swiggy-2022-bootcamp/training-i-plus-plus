package model

type Patient struct {
	Id 					string
	User
	DoctorAssignedId	string
	IsDischarged		bool
	RoomAllocated		string
}