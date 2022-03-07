package db

import (
	"errors"
	"shopping-app/user/domain"
	"time"
)

type userRepositoryDB struct {
	userList []User
}

func NewUserRepositoryDB(uList []User) domain.UserRepositoryDB {
	return &userRepositoryDB{
		userList: uList,
	}
}

func (udb *userRepositoryDB) Save(u domain.User) (domain.User, error) {

	user := NewUser(
		u.Email(),
		u.Password(),
		u.Name(),
		u.Address(),
		u.Zipcode(),
		u.MobileNo(),
		u.Role(),
	)
	user.SetCreatedAt(time.Now())
	user.SetUpdatedAt(time.Now())

	udb.userList = append(udb.userList, *user)
	return u, nil
}

func (udb *userRepositoryDB) FindUserByEmail(email string) (*domain.User, error) {

	for _, user := range udb.userList {
		if user.email == email {
			domainUser := domain.NewUser(user.email, user.password, user.name, user.address, user.zipcode, user.mobileNo, user.role)
			return domainUser, nil
		}
	}
	return nil, errors.New("user does not exist")
}
