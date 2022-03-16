package domain_test

import (
	"errors"
	"panem/domain"
	"panem/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

var mockUserRepo = mocks.UserRepository{}
var userService = domain.NewUserService(&mockUserRepo)

func TestShouldReturnNewUserService(t *testing.T) {
	userService := domain.NewUserService(nil)
	assert.NotNil(t, userService)
}

func TestShouldCreateNewUser(t *testing.T) {
	firstName := "Murtaza"
	lastName := "Sadriwala"
	phone := "9900887766"
	email := "murtaza896@gmail.com"
	username := "murtaza896"
	password := "Pass!23"
	role := domain.Admin

	user := domain.NewUser(firstName, lastName, username, phone, email, password, role)
	mockUserRepo.On("Save", *user).Return(*user, nil)
	userService.CreateUser(firstName, lastName, username, phone, email, password, role)
	mockUserRepo.AssertNumberOfCalls(t, "Save", 1)
}

func TestShouldDeleteUserByUsername(t *testing.T) {
	username := "testUsername"
	mockUserRepo.On("DeleteUserByUsername", username).Return(true, nil)
	mockUserRepo.On("FindByUsername", username).Return(nil, nil)

	res, _ := userService.DeleteUserByUsername(username)
	assert.Equal(t, true, res)
}

func TestShouldNotDeleteUserByUsernameUponInvalidUsername(t *testing.T) {
	username := "invalidUsername"
	errMessage := "some error"
	mockUserRepo.On("DeleteUserByUsername", username).Return(false, nil)
	mockUserRepo.On("FindByUsername", username).Return(nil, errors.New(errMessage))

	res, err := userService.DeleteUserByUsername(username)
	assert.Error(t, err, errMessage)
	assert.Equal(t, false, res)
}
