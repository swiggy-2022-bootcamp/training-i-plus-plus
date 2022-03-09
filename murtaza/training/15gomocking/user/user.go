package user

import "github.com/nishant-nimbare/ipp/go_mocking/doer"

type User struct {
	Doer doer.Doer
}

func (u *User) Use() error {
	return u.Doer.DoSomething(5, "From User")
}
