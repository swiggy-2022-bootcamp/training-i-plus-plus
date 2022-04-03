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
	hashedPassword, _ := domain.HashPassword("Pass!23")
	role := domain.Admin

	user := domain.NewUser(firstName, lastName, username, phone, email, hashedPassword, role)
	user.Id = 1
	mockUserRepo.On("FindUserByUsername", username).Return(user, nil)
	res, _ := authService.AuthenticateUser(username, password)
	assert.NotNil(t, res)
	assert.True(t, len(res) > 0)
}

func TestShouldParseValidAuthToken(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MCwicm9sZSI6MX0.EbRVzFSPslw_bhnchRWTGPRNlJnsqWXsbXrirjvA9ss"

	actualResponse, actualRole, err := authService.ParseAuthToken(tokenString)
	expectedResponse := 0
	expectedRole := domain.Seller
	assert.Nil(t, err)
	assert.Equal(t, expectedResponse, actualResponse)
	assert.Equal(t, expectedRole, actualRole)
}

func TestShouldNotParseAuthTokenWithInvalidSecret(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MH0.WOK8U6-k2NvDSyHMLQM2YJN3whV4LU-OtL2FiRDpZWM"

	_, _, err := authService.ParseAuthToken(tokenString)
	assert.NotNil(t, err)
}
