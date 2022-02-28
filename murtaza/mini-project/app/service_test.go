package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var userService = NewService()

func TestAuthenticateUserShouldReturnTokenOnValidCredentials(t *testing.T) {
	actualResponse, err := userService.AuthenticateUser("murtaza896", "Pass!23")
	assert.Nil(t, err)
	assert.NotNil(t, actualResponse)
	assert.True(t, len(actualResponse) > 0)
}

func TestShouldParseValidAuthToken(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MH0.hRNAVCB1LFZSCYGW-S8GMZXl1AVAmN2nFZoysNJUBrY"

	actualResponse, err := userService.ParseAuthToken(tokenString)
	expectedResponse := float64(0)
	assert.Nil(t, err)
	assert.Equal(t, expectedResponse, actualResponse)
}

func TestShouldNotParseAuthTokenWithInvalidSecret(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MH0.WOK8U6-k2NvDSyHMLQM2YJN3whV4LU-OtL2FiRDpZWM"

	_, err := userService.ParseAuthToken(tokenString)
	assert.NotNil(t, err)
}
