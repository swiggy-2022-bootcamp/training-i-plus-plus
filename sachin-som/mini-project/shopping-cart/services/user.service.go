package services

import "github.com/sachinsom93/shopping-cart/models"

type UserServices interface {
	CreateUser(*models.User) error
	GetUser(*string) (*models.User, error)
	GetAllUser() ([]*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(*string) error
}
