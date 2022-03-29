package domain

type User struct {
	Email    string
	Password string
	Name     string
	Address  string
	Zipcode  int32
	MobileNo string
	Role     string
}

func NewUser(email string, password string, name string, address string, zipcode int32, mobileNo string, role string) *User {
	return &User{
		Email:    email,
		Password: password,
		Name:     name,
		Address:  address,
		Zipcode:  zipcode,
		MobileNo: mobileNo,
		Role:     role,
	}
}

// func (u *User) Email() string {
// 	return u.email
// }

// func (u *User) Password() string {
// 	return u.password
// }

// func (u *User) Name() string {
// 	return u.name
// }

// func (u *User) Address() string {
// 	return u.address
// }

// func (u *User) Zipcode() int32 {
// 	return u.zipcode
// }

// func (u *User) MobileNo() string {
// 	return u.mobileNo
// }

// func (u *User) Role() string {
// 	return u.role
// }

func (u *User) SetEmail(email string) {
	u.Email = email
}

func (u *User) SetPassword(password string) {
	u.Password = password
}

func (u *User) SetName(name string) {
	u.Name = name
}

func (u *User) SetAddress(address string) {
	u.Address = address
}

func (u *User) SetZipcode(zipcode int32) {
	u.Zipcode = zipcode
}

func (u *User) SetMobileNo(mobileNo string) {
	u.MobileNo = mobileNo
}

func (u *User) SetRole(role string) {
	u.Role = role
}

//go:generate mockgen -destination=../mocks/domain/mockUser.go -package=domain users/domain User
type UserRepositoryDB interface {
	Save(User) (*User, error)
	FindUserByEmail(string) (*User, error)
	UpdateUser(User) (*User, error)
	DeleteUserByEmail(string) error
}
