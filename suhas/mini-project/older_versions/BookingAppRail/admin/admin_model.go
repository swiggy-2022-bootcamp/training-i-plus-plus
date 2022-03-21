package admin

type Admin struct {
	Name   string
	Userid int
}

var Admin_array = []Admin {
	{"admin1",101},
	{"admin2",102},
	{"admin3",103},
}

func GetAdminArray () []Admin {
	return Admin_array
}