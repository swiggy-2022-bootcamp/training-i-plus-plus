package domain

import "errors"

type UserService interface {
	Register(User) (*User, error)
	Login(string, string) (string, error)
	FindUserByEmail(string) (*User, error)
	VerifyCredentials(string, string) (bool, error)
	VerifySavedToken(string, string) (bool, error)
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

	_, err := usvc.FindUserByEmail(user.email)
	if err != nil {
		return nil, errors.New("user already exists")
	}
	u, err := usvc.userDB.Save(user)
	return &u, err
}

func (usvc *DefaultUserService) Login(email string, password string) (string, error) {

	isValid, err := usvc.VerifyCredentials(email, password)
	if err != nil {
		return "", errors.New("user already exists")
	}
	if isValid {
		token := "$" + email + "$" + password + "$"
		return token, nil
	}
	return "", errors.New("invalid user credentials")
}

func (usvc *DefaultUserService) VerifyCredentials(email string, password string) (bool, error) {
	user, err := usvc.FindUserByEmail(email)
	if err != nil {
		return false, err
	}
	if password == user.Password() {
		return true, nil
	}
	return false, errors.New("invalid credentials")
}

func (usvc *DefaultUserService) FindUserByEmail(email string) (*User, error) {

	user, err := usvc.userDB.FindUserByEmail(email)
	return user, err
}

func (usvc *DefaultUserService) VerifySavedToken(email string, token string) (bool, error) {

	user, err := usvc.FindUserByEmail(email)
	if err != nil {
		return false, err
	}
	actualToken, err := usvc.Login(email, user.Password())
	if err != nil {
		return false, err
	}
	return actualToken == token, nil
}
