package domain

type User struct {
	email    string
	password string
	name     string
	address  string
	zipcode  int32
	mobileNo string
	role     string
}

func NewUser(email string, password string, name string, address string, zipcode int32, mobileNo string, role string) *User {
	return &User{
		email:    email,
		password: password,
		name:     name,
		address:  address,
		zipcode:  zipcode,
		mobileNo: mobileNo,
		role:     role,
	}
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Password() string {
	return u.password
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Address() string {
	return u.address
}

func (u *User) Zipcode() int32 {
	return u.zipcode
}

func (u *User) MobileNo() string {
	return u.mobileNo
}

func (u *User) Role() string {
	return u.role
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u *User) SetPassword(password string) {
	u.password = password
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) SetAddress(address string) {
	u.address = address
}

func (u *User) SetZipcode(zipcode int32) {
	u.zipcode = zipcode
}

func (u *User) SetMobileNo(mobileNo string) {
	u.mobileNo = mobileNo
}

func (u *User) SetRole(role string) {
	u.role = role
}

type UserRepositoryDB interface {
	Save(User) (*User, error)
	FindUserByEmail(string) (*User, error)
	UpdateUser(User) (*User, error)
}
