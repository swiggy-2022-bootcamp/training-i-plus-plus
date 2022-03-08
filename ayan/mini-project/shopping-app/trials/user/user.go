package user

type UserService interface {
	NewUser() (User, error)
	RegisterUser(email string, password string, name string, address string, pincode int32, mobileNo int64) (string, error)
	LoginUser(email string, password string) (string, error)
	// UpdateUser() (User, error)
	
	ChangePassword(email string, oldPassword string, newPassword string)
}

type User struct {
	email    string
	password string
	name     string
	address  string
	pincode  int32
	mobileNo int64
}

func (user *User) validateUser() bool {

	isValid := true

	if len(user.email) == 0 || len(user.password) == 0 || len(user.name) == 0 {
		isValid = false
	}

	return isValid
}

func (user *UserService) RegisterUser(email string, password string, name string, address string, pincode int32, mobileNo int64) (string, error) {

	user := 
}
