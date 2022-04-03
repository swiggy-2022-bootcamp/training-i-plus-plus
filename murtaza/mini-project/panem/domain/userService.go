package domain

import (
	"panem/utils/errs"
)

type UserService interface {
	CreateUserInMongo(string, string, string, string, string, string, Role) (User, *errs.AppError)
	GetMongoUserByUserId(int) (*User, *errs.AppError)
	DeleteUserByUserId(int) *errs.AppError
	UpdateUser(User) (*User, *errs.AppError)
}

type service struct {
	userMongoRepository UserMongoRepository
}

func (s service) CreateUserInMongo(firstName, lastName, username, phone, email, password string, role Role) (User, *errs.AppError) {
	user := NewUser(firstName, lastName, username, phone, email, password, role)
	persistedUser, err := s.userMongoRepository.InsertUser(*user)
	if err != nil {
		return User{}, err
	}
	return persistedUser, nil
}

func (s service) GetMongoUserByUserId(userId int) (*User, *errs.AppError) {
	res, err := s.userMongoRepository.FindUserById(userId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s service) DeleteUserByUserId(userId int) *errs.AppError {
	err := s.userMongoRepository.DeleteUserByUserId(userId)
	if err != nil {
		return err
	}
	return nil
}

func (s service) UpdateUser(user User) (*User, *errs.AppError) {
	res, err := s.userMongoRepository.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewUserService(userMongoRepository UserMongoRepository) UserService {
	return &service{
		userMongoRepository: userMongoRepository,
	}
}
