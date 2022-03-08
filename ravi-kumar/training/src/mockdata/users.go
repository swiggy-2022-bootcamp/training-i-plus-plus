package mockdata

type User struct {
	Fullname string
	UserName string
	Password string
	Address  string
}

func GetAllUsers() []User {
	var allUsers = []User{
		{
			Fullname: "Ravi Kumar",
			UserName: "ravi",
			Password: "Password",
			Address:  "Bangalore",
		},
		{
			Fullname: "Ajay",
			UserName: "ajay99",
			Password: "Password",
			Address:  "Delhi",
		},
	}
	return allUsers
}

func Authenticate(UserName string, Password string) bool {
	users := GetAllUsers()
	for i := 0; i < len(users); i++ {
		if users[i].UserName == UserName && users[i].Password == Password {
			return true
		}
	}
	return false
}
