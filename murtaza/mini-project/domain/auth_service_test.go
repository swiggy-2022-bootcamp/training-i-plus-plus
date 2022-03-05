package domain_test

import (
	"mini-project/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

var authService = domain.NewAuthService(&mockUserRepo)

func TestShouldGenereateJWTToken(t *testing.T) {
	firstName := "Murtaza"
	lastName := "Sadriwala"
	phone := "9900887766"
	email := "murtaza896@gmail.com"
	username := "murtaza896"
	password := "Pass!23"
	role := domain.Admin

	user := domain.NewUser(firstName, lastName, username, phone, email, password, role)
	user.SetId(1)
	mockUserRepo.On("FindByUsername", username).Return(user, nil)
	res, _ := authService.AuthenticateUser(username, password)
	assert.NotNil(t, res)
}

func TestShouldParseAuthToken(t *testing.T) {
	firstName := "Murtaza"
	lastName := "Sadriwala"
	phone := "9900887766"
	email := "murtaza896@gmail.com"
	username := "murtaza896"
	password := "Pass!23"
	role := domain.Admin

	user := domain.NewUser(firstName, lastName, username, phone, email, password, role)
	user.SetId(1)
	mockUserRepo.On("FindByUsername", username).Return(user, nil)
	authToken, _ := authService.AuthenticateUser(username, password)
	actualId, actualRole, _ := authService.ParseAuthToken(authToken)
	assert.Equal(t, 1, actualId)
	assert.Equal(t, role, actualRole)
}
