package admin

 import (
 	"booking_app/train"
	"booking_app/storage/handler"
 )

type Admin struct {
	Name   string
	Userid int
}

func (admin Admin) CheckAdmin() bool {
   //memory
   for _,v := range(storage.Admin_array) {
	   if v.Userid == admin.Userid {
		   return true
	   }
   }
   return false
}

func (admin Admin) AddTrainAdmin(train train.Train) bool {
	//memory
	storage.Admin_array = append(storage.Admin_array,train)
	return true
}

func (admin Admin) GetTrains() []train.Train{
	//memory
	return storage.Train_array
}

func (admin Admin) AddTicketAdmin(ticket TicketAvailable) bool {
	//memory
	return true
}

  
