package domain

import (
	"user/utils"
	"user/utils/errs"
)

type UserService interface {
	Register(User) (*User, *errs.AppError)
	Login(string, string) (string, *errs.AppError)
	FindByEmail(string) (*User, *errs.AppError)
	VerifyCredentials(string, string) (bool, *errs.AppError)
	VerifyToken(string, string, string) (bool, *errs.AppError)
	Update(User) (*User, *errs.AppError)
	DeleteByEmail(string) (*User, *errs.AppError)
}

type DefaultUserService struct {
	UserDB UserRepositoryDB
}

func NewUserService(userDB UserRepositoryDB) UserService {
	return &DefaultUserService{
		UserDB: userDB,
	}
}

func (usvc DefaultUserService) Register(user User) (*User, *errs.AppError) {

	_, err := usvc.FindByEmail(user.Email)
	if err == nil {
		return nil, errs.NewValidationError("user already exists")
	}
	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hash
	u, err := usvc.UserDB.Save(user)
	return u, err
}

func (usvc DefaultUserService) Login(email string, password string) (string, *errs.AppError) {

	isValid, err := usvc.VerifyCredentials(email, password)
	if err != nil {
		return "", err
	}
	if isValid {
		user, err := usvc.FindByEmail(email)
		if err != nil {
			return "", err
		}
		token, err := utils.GenerateJWT(email, user.Role)
		return token, err
	}
	return "", errs.NewAuthenticationError("invalid credentials")
}

func (usvc DefaultUserService) VerifyCredentials(email string, password string) (bool, *errs.AppError) {
	user, err := usvc.FindByEmail(email)
	if err != nil {
		return false, err
	}
	if utils.CheckPasswordHash(password, user.Password) {
		return true, nil
	}
	return false, errs.NewAuthenticationError("invalid credentials")
}

func (usvc *DefaultUserService) FindByEmail(email string) (*User, *errs.AppError) {

	user, err := usvc.UserDB.FetchUserByEmail(email)
	return user, err
}

func (usvc DefaultUserService) VerifyToken(email string, role string, token string) (bool, *errs.AppError) {

	_, err := usvc.FindByEmail(email)
	if err != nil {
		return false, err
	}

	tokenEmail, tokenRole, err := utils.ParseAuthToken(token)
	if err != nil {
		return false, err
	}
	return tokenEmail == email && tokenRole == role, nil
}

func (usvc DefaultUserService) Update(user User) (*User, *errs.AppError) {

	_, err := usvc.FindByEmail(user.Email)
	if err != nil {
		return nil, err
	}
	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hash
	u, err := usvc.UserDB.UpdateUser(user)
	return u, err
}

func (usvc DefaultUserService) DeleteByEmail(email string) (*User, *errs.AppError) {

	u, err := usvc.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	err = usvc.UserDB.DeleteUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return u, nil
}
