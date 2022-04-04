package domain

import (
	"user/utils/errs"
)

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

type UserRepositoryDB interface {
	Save(User) (*User, *errs.AppError)
	FetchUserByEmail(string) (*User, *errs.AppError)
	UpdateUser(User) (*User, *errs.AppError)
	DeleteUserByEmail(string) *errs.AppError
}
