package db

import (
	"errors"
	"time"
	"user/domain"
)

type userRepositoryListDB struct {
	userList []User
}

func NewUserRepositoryListDB(uList []User) domain.UserRepositoryDB {
	return &userRepositoryListDB{
		userList: uList,
	}
}

func (udb *userRepositoryListDB) Save(u domain.User) (*domain.User, error) {

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
	return &u, nil
}

func (udb *userRepositoryListDB) FindUserByEmail(email string) (*domain.User, error) {

	for _, user := range udb.userList {
		if user.email == email {
			domainUser := domain.NewUser(user.email, user.password, user.name, user.address, user.zipcode, user.mobileNo, user.role)
			return domainUser, nil
		}
	}
	return nil, errors.New("user does not exist")
}

func (udb userRepositoryListDB) UpdateUser(u domain.User) (*domain.User, error) {
	// NOOP
	return &u, nil
}
