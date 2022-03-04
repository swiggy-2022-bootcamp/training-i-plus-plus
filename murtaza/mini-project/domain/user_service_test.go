package domain_test

import (
	"mini-project/domain"
	"mini-project/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnNewUserService(t *testing.T) {
	userService := domain.NewUserService(nil)
	assert.NotNil(t, userService)
}

func TestShouldCreateNewUser(t *testing.T) {
	mockUserRepo := mocks.UserRepository{}
	userService := domain.NewUserService(&mockUserRepo)

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
