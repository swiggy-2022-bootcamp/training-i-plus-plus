package service

import (
	repository "User-Service/Repository"
	"User-Service/Repository/mocks"
	"User-Service/errors"
	mockdata "User-Service/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

var mongoDAO *mocks.MongoDAO = &mocks.MongoDAO{}
var userService UserService

func TestShouldAvertInvalidLogin(t *testing.T) {
	userName := "InvalidUserName"
	password := "InvalidPassword"

	logInDTO := mockdata.LogInDTO{UserName: userName, Password: password}

	//when there's no need to mock : init userService's mongoDAO with concrete implementation of IMongoDAO - i.e MongoDAO
	userService := InitUserService(&repository.MongoDAO{})
	_, err := userService.LogInUser(logInDTO)

	assert.EqualError(t, err, errors.UnauthorizedError().ErrorMessage)
}

func TestShouldApproveValidLogin(t *testing.T) {
	userName := "InvalidUserName"
	password := "InvalidPassword"

	logInDTO := mockdata.LogInDTO{UserName: userName, Password: password}

	userService := InitUserService(&repository.MongoDAO{})
	_, err := userService.LogInUser(logInDTO)

	assert.EqualError(t, err, errors.UnauthorizedError().ErrorMessage)
}

func TestShouldCreateUser(t *testing.T) {
	newUser := mockdata.User{
		Fullname: "sarath",
		UserName: "sharath1011",
		Password: "Password",
		Address:  "Chennai",
		Role:     2,
	}

	mongoDAO.On("MongoCreateUser", newUser).Return("624865bc030ba951de8956bd")

	//when there's need to mock : init userService's mongoDAO with mockery's implementation of IMongoDAO - i.e mongoDAO
	userService := InitUserService(mongoDAO)
	insertedId, _, err := userService.CreateUser(newUser)

	assert.Nil(t, err)
	assert.Equal(t, insertedId, "624865bc030ba951de8956bd")
}

func TestAssuresGetAllUsersShouldReturnAtleastOneUser(t *testing.T) {
	userService := InitUserService(&repository.MongoDAO{})
	assert.Greater(t, len(userService.GetAllUsers()), 0)
}

func TestShouldReturnErrorWhenUserIdIsMalformed(t *testing.T) {
	updatedUser := mockdata.User{
		Fullname: "sarath",
		UserName: "sharath1011",
		Password: "Password",
		Address:  "Chennai",
		Role:     2,
	}

	malformedUserId := "624865bc030ba951de8956bderred23"

	userService := InitUserService(&repository.MongoDAO{})

	_, err := userService.UpdateUserById(malformedUserId, updatedUser)
	assert.EqualError(t, err, errors.MalformedIdError().ErrorMessage)

	_, err = userService.GetUserById(malformedUserId)
	assert.EqualError(t, err, errors.MalformedIdError().ErrorMessage)

	_, err = userService.DeleteUserbyId(malformedUserId)
	assert.EqualError(t, err, errors.MalformedIdError().ErrorMessage)
}

func TestShouldReturnErrorWhenUserIdIsNotFound(t *testing.T) {
	updatedUser := mockdata.User{
		Fullname: "sarath",
		UserName: "sharath1011",
		Password: "Password",
		Address:  "Chennai",
		Role:     2,
	}

	userId := "624865bc030ba951de8956ba"

	userService := InitUserService(&repository.MongoDAO{})

	mongoDAO.On("MongoGetUserById", userId).Return(nil, errors.IdNotFoundError())
	mongoDAO.On("MongoUpdateUserById", userId, updatedUser).Return(nil, errors.IdNotFoundError())
	mongoDAO.On("MongoDeleteUserById", userId).Return(nil, errors.IdNotFoundError())

	_, err := userService.GetUserById(userId)
	assert.EqualError(t, err, errors.IdNotFoundError().ErrorMessage)

	_, err = userService.UpdateUserById(userId, updatedUser)
	assert.EqualError(t, err, errors.IdNotFoundError().ErrorMessage)

	_, err = userService.DeleteUserbyId(userId)
	assert.EqualError(t, err, errors.IdNotFoundError().ErrorMessage)
}
