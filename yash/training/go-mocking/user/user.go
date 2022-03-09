package user
import "github.com/sgreben/testing-with-gomock/doer"

type User struct{
	Doer doer.Doer
}
func(u *User) use() error{
	return u.Doer.DoSomething(123, "go mock")
}