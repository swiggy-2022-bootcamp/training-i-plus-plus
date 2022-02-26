package user

import "github.com/sgreben/testing-with-gomock/doer"

type User struct {
	Doer doer.Doer
}

func (u *User) User() error {
	return u.Doer.DoSomething(123, "Hello GoMock")
}
