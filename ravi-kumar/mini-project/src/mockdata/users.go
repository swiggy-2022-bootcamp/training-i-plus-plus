package mockdata

type User struct {
	UserId   int
	Fullname string
	UserName string
	Password string
	Address  string
}

func GetAllUsers() []User {
	var allUsers = []User{
		{
			UserId:   1,
			Fullname: "Ravi Kumar",
			UserName: "ravi",
			Password: "Password",
			Address:  "Bangalore",
		},
		{
			UserId:   2,
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
