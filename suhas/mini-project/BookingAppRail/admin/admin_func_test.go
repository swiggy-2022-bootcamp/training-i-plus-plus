package admin

 import (
// 	"BookingApp/train"
"testing"
"BookingApp/ticket"
)

type testCaseCheck struct {
    arg1 Admin
    exist bool
}

var testCasesC []testCaseCheck = []testCaseCheck{
	{arg1:Admin{},exist:false},
	{arg1:Admin{"",101},exist:false},
	{arg1:Admin{"admin14",10},exist:false},
	{arg1:Admin{"admin3",103},exist:true},
}
	
func TestCheckAdmin(t *testing.T) {
	for _,v := range(testCasesC) {
		res:= v.arg1.CheckAdmin()
		if res != v.exist {
			t.Error(v)
			t.Error(res)
			t.Errorf("Expected'%t', but got '%t'", v.exist, res)
		}
	}
}


var testCasesAdd []testCaseCheck = []testCaseCheck {
	{arg1:Admin{"admin5",105},exist:true},
	{arg1:Admin{"admin6",106},exist:true},
	{arg1:Admin{"admin7",107},exist:true},
	{arg1:Admin{"admin8",108},exist:true},
}


func TestAddAdmin(t *testing.T) {
	for _,v := range(testCasesAdd) {
		res:= v.arg1.AddAdmin()
		if res != v.exist {
			t.Error(v)
			t.Error(res)
			t.Errorf("Expected'%t', but got '%t'", v.exist, res)
		}
	}
}

type testCasesAddAdminTicket struct {
    arg1 ticket.TicketAvailable
    exist bool
}

var testCaseAddAdminTicket []testCasesAddAdminTicket = []ticket.testCasesAddAdminTicket{
	{arg1:{},true},
	{arg1:{},true},
	{arg1:{},true},
	{arg1:{},true},
}