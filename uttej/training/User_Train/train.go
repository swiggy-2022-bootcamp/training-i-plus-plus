package train

import "fmt"

type Module interface {
	Details() string
}

type Train struct {
	name  string
	seats int32
}

func (t Train) Details() string {
	str1 := fmt.Sprintf("%s has a capacity of %d seats", t.name, t.seats)
	return str1
}

type User struct {
	userName string
	age      int32
	gender   string
}

func (u User) Details() string {
	str1 := fmt.Sprintf("Hey %s, you are %d years old and %s", u.userName, u.age, u.gender)
	return str1
}

func Info(m Module) string {
	return m.Details()
}
