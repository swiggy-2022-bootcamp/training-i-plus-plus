package admin

 import (
// 	"BookingApp/train"
"BookingApp/ticket"
 )

func (admin Admin) CheckAdmin() bool {
	//memory
	for _,v := range(Admin_array) {
		if v.Userid == admin.Userid && v.Name == admin.Name {
			return true
		}
	}
	return false
}

 
func (adm Admin) AddAdmin() bool {
	//memory
	Admin_array = append(Admin_array,adm)
	return true
}
 
func (admin Admin) AddTicketAdmin(tk ticket.TicketAvailable) bool {
	//memory
	if ticket.IsTicketAvialable(tk.Train_id) == false {
		return false
	}
	ticket.Ticket_array = append(ticket.Ticket_array,tk)
	return true
}

func DeleteAdmin(ad Admin) bool{
	for i,v := range(Admin_array) {
		if v.Userid == ad.Userid && v.Name == ad.Name {
			Admin_array = append(Admin_array[:i], Admin_array[i+1:]...)
			return true
		}
	}
	return false
}

func (admin Admin) UpdateAdmin(newadmin Admin) {
	DeleteAdmin(admin)
	newadmin.AddAdmin()
}