package main

import (
	"errors"
)

type Service struct {
	userList    []User
	tokenMap    map[string]string
	productList []Product
}

type User struct {
	email    string
	password string
	name     string
	address  string
	pincode  int32
	mobileNo int64
}

type Product struct {
	name     string
	category string
	price    string
	quantity string
}

func (svc *Service) FindUserByEmail(email string) (User, bool) {

	for _, user := range svc.userList {
		if user.email == email {
			return user, true
		}
	}
	return User{}, false
}

func (svc *Service) AddUser(user User) {

	svc.userList = append(svc.userList, user)
}

func (svc *Service) RegisterUser(user User) (*User, error) {

	_, isUserPresent := svc.FindUserByEmail(user.email)
	if isUserPresent {
		return nil, errors.New("user already exists")
	}
	svc.AddUser(user)
	return &user, nil
}

func (svc *Service) VerifyUserCredentials(email string, password string) bool {

	user, ok := svc.FindUserByEmail(email)
	if ok {
		return password == user.password
	}
	return false
}

func (svc *Service) LoginUser(email string, password string) (string, error) {

	isValid := svc.VerifyUserCredentials(email, password)
	if isValid {
		token := "$" + email + "$" + password + "$"
		svc.tokenMap[email] = token
		return token, nil
	}
	return "", errors.New("invalid user credentials")
}

func (svc *Service) VerifySavedUserToken(email string, token string) bool {

	savedToken, ok := svc.tokenMap[email]
	if ok {
		return token == savedToken
	}
	return false
}

func (svc *Service) LogoutUser(email string) (bool, error) {

	_, isUserPresent := svc.FindUserByEmail(email)
	if isUserPresent {

		_, ok := svc.tokenMap[email]
		if ok {
			delete(svc.tokenMap, email)
			return true, nil
		}

		return false, errors.New("user not logged in")
	}
	return false, errors.New("invalid user email")
}
