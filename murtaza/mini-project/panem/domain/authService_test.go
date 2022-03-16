package domain_test

import (
	"github.com/stretchr/testify/assert"
	"panem/domain"
	"testing"
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

func TestAuthenticateUserShouldReturnTokenOnValidCredentials(t *testing.T) {
	// var userService = NewService()
	// actualResponse, err := userService.AuthenticateUser("murtaza896", "Pass!23")
	// assert.Nil(t, err)
	// assert.NotNil(t, actualResponse)
	// assert.True(t, len(actualResponse) > 0)
}

func TestShouldParseValidAuthToken(t *testing.T) {
	// var userService = NewService()
	// tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MCwicm9sZSI6MX0.EbRVzFSPslw_bhnchRWTGPRNlJnsqWXsbXrirjvA9ss"

	// actualResponse, actualRole, err := userService.ParseAuthToken(tokenString)
	// expectedResponse := int(0)
	// expectedRole := Admin
	// assert.Nil(t, err)
	// assert.Equal(t, expectedResponse, actualResponse)
	// assert.Equal(t, expectedRole, actualRole)
}

func TestShouldNotParseAuthTokenWithInvalidSecret(t *testing.T) {
	// var userService = NewService()
	// tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MH0.WOK8U6-k2NvDSyHMLQM2YJN3whV4LU-OtL2FiRDpZWM"

	// _, _, err := userService.ParseAuthToken(tokenString)
	// assert.NotNil(t, err)
}
