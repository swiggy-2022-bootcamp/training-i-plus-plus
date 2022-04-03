package db

import (
	"sample.akash.com/model"
)

type CustomerRepository interface {
	Connect()
	FindOneWithUsername(username string) *model.User
	FindAll() []model.User
	SaveUser(user model.User)
	DeleteUser(username string) bool
	FindAndUpdate(user model.User) bool
}
