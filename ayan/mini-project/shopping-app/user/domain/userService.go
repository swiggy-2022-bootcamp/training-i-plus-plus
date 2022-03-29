package domain

import (
	"errors"
	"fmt"

	"user/utils"
)

//go:generate D:/Users/ayandut/go/bin/mockgen.exe -destination=../mocks/domain/mockUserService.go -package=domain users/domain UserService
type UserService interface {
	Register(User) (*User, error)
	Login(string, string) (string, error)
	FindUserByEmail(string) (*User, error)
	VerifyCredentials(string, string) (bool, error)
	VerifyToken(string, string) (bool, error)
	UpdateUser(User) (*User, error)
	DeleteUserByEmail(string) (*User, error)
}

type DefaultUserService struct {
	userDB UserRepositoryDB
}

func NewUserService(userDB UserRepositoryDB) UserService {
	return &DefaultUserService{
		userDB: userDB,
	}
}

func (usvc *DefaultUserService) Register(user User) (*User, error) {

	_, err := usvc.FindUserByEmail(user.Email)
	fmt.Println("(usvc *DefaultUserService) Register : ", err)
	if err == nil {
		fmt.Println("(usvc *DefaultUserService) Register : err is nil")
		return nil, errors.New("user already exists")
	}
	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hash
	u, err := usvc.userDB.Save(user)
	return u, err
}

func (usvc *DefaultUserService) Login(email string, password string) (string, error) {

	isValid, err := usvc.VerifyCredentials(email, password)
	if err != nil {
		return "", err
	}
	if isValid {
		hash, _ := utils.HashPassword(password)
		token := "$" + email + "$" + hash + "$"
		return token, nil
	}
	return "", errors.New("invalid user credentials")
}

func (usvc *DefaultUserService) VerifyCredentials(email string, password string) (bool, error) {
	user, err := usvc.FindUserByEmail(email)
	if err != nil {
		return false, err
	}
	if utils.CheckPasswordHash(password, user.Password) {
		return true, nil
	}
	return false, errors.New("invalid credentials")
}

func (usvc *DefaultUserService) FindUserByEmail(email string) (*User, error) {

	user, err := usvc.userDB.FindUserByEmail(email)
	fmt.Println("(usvc *DefaultUserService) FindUserByEmail : ", user, err)
	return user, err
}

func (usvc *DefaultUserService) VerifyToken(email string, token string) (bool, error) {

	user, err := usvc.FindUserByEmail(email)
	if err != nil {
		return false, err
	}
	actualToken, err := usvc.Login(email, user.Password)
	if err != nil {
		return false, err
	}
	return actualToken == token, nil
}

func (usvc *DefaultUserService) UpdateUser(user User) (*User, error) {

	_, err := usvc.FindUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}
	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hash
	u, err := usvc.userDB.UpdateUser(user)
	return u, err
}

func (usvc *DefaultUserService) DeleteUserByEmail(email string) (*User, error) {

	u, err := usvc.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}
	err = usvc.userDB.DeleteUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return u, nil
}
