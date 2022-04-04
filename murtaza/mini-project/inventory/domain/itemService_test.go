package domain_test

//
//import (
//	"github.com/stretchr/testify/mock"
//	"panem/domain"
//	"panem/mocks"
//	"panem/utils/errs"
//	"testing"
//
//	"github.com/stretchr/testify/assert"
//)
//
//var mockUserRepo = mocks.UserMongoRepository{}
//var userService = domain.NewUserService(&mockUserRepo)
//
//func TestShouldReturnNewUserService(t *testing.T) {
//	userService := domain.NewUserService(nil)
//	assert.NotNil(t, userService)
//}
//
//func TestShouldCreateNewUser(t *testing.T) {
//	firstName := "Murtaza"
//	lastName := "Sadriwala"
//	phone := "9900887766"
//	email := "murtaza896@gmail.com"
//	username := "murtaza896"
//	password, _ := domain.HashPassword("Pass!23")
//	role := domain.Admin
//
//	user := domain.NewUser(firstName, lastName, username, phone, email, password, role)
//	mockUserRepo.On("InsertItem", mock.Anything).Return(*user, nil)
//	userService.CreateUserInMongo(firstName, lastName, username, phone, email, password, role)
//	mockUserRepo.AssertNumberOfCalls(t, "InsertItem", 1)
//}
//
//func TestShouldDeleteUserByUserId(t *testing.T) {
//	userId := 1
//	mockUserRepo.On("DeleteUserByUserId", userId).Return(nil)
//
//	var err = userService.DeleteUserByUserId(userId)
//	assert.Nil(t, err)
//}
//
//func TestShouldNotDeleteUserByUserIdUponInvalidUserId(t *testing.T) {
//	userId := -99
//	errMessage := "some error"
//	mockUserRepo.On("DeleteUserByUserId", userId).Return(errs.NewUnexpectedError(errMessage))
//
//	err := userService.DeleteUserByUserId(userId)
//	assert.Error(t, err.Error(), errMessage)
//}
