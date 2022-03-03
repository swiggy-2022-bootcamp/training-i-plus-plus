package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var mockUser = User{
	firstName: "aryan",
	lastName:  "singh",
	username:  "aryan123",
	password:  "passw0rd",
	email:     "aryan.singh@gmail.com",
	phone:     "9900991234",
	role:      Seller,
}

func TestAuthenticateUserShouldReturnTokenOnValidCredentials(t *testing.T) {
	var userService = NewService()
	actualResponse, err := userService.AuthenticateUser("murtaza896", "Pass!23")
	assert.Nil(t, err)
	assert.NotNil(t, actualResponse)
	assert.True(t, len(actualResponse) > 0)
}

func TestShouldParseValidAuthToken(t *testing.T) {
	var userService = NewService()
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MCwicm9sZSI6MX0.EbRVzFSPslw_bhnchRWTGPRNlJnsqWXsbXrirjvA9ss"

	actualResponse, actualRole, err := userService.ParseAuthToken(tokenString)
	expectedResponse := int(0)
	expectedRole := Admin
	assert.Nil(t, err)
	assert.Equal(t, expectedResponse, actualResponse)
	assert.Equal(t, expectedRole, actualRole)
}

func TestShouldNotParseAuthTokenWithInvalidSecret(t *testing.T) {
	var userService = NewService()
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MH0.WOK8U6-k2NvDSyHMLQM2YJN3whV4LU-OtL2FiRDpZWM"

	_, _, err := userService.ParseAuthToken(tokenString)
	assert.NotNil(t, err)
}

func TestShouldAddUser(t *testing.T) {
	var userService = NewService()
	newUser := User{
		firstName: "aryan",
		lastName:  "singh",
		username:  "aryan123",
		password:  "passw0rd",
		email:     "aryan.singh@gmail.com",
		phone:     "9900991234",
		role:      Seller,
	}

	expectedLength := len(userService.GetAllUsers()) + 1
	res, _ := userService.AddUser(newUser)
	actualLength := len(userService.GetAllUsers())

	assert.Equal(t, true, res)
	assert.Equal(t, expectedLength, actualLength)
}

func TestShouldDeleteUserByUsername(t *testing.T) {
	var userService = NewService()
	var username string = "murtaza896"

	expectedLength := len(userService.GetAllUsers()) - 1
	res, _ := userService.DeleteUserByUsername(username)
	actualLength := len(userService.GetAllUsers())

	assert.Equal(t, true, res)
	assert.Equal(t, expectedLength, actualLength)
}

func TestShouldReturnUserByUsername(t *testing.T) {
	var userService = NewService()
	username := "murtaza896"
	user, err := userService.GetUserByUsername(username)
	assert.Equal(t, username, user.GetUsername())
	assert.NotNil(t, user)
	assert.Nil(t, err)
}

func TestShouldReturnErrOnInvalidUsername(t *testing.T) {
	var userService = NewService()
	username := "murtaza89612"
	user, err := userService.GetUserByUsername(username)
	assert.Nil(t, user)
	assert.NotNil(t, err)
}

func TestShouldReturnUserByEmail(t *testing.T) {
	var userService = NewService()
	email := "murtaza896@gmail.com"
	user, err := userService.GetUserByEmail(email)
	assert.Equal(t, email, user.email)
	assert.NotNil(t, user)
	assert.Nil(t, err)
}

func TestShouldReturnErrOnInvalidEmail(t *testing.T) {
	var userService = NewService()
	email := "fake.email"
	user, err := userService.GetUserByEmail(email)
	assert.Nil(t, user)
	assert.NotNil(t, err)
}

func TestShouldUpdateEmail(t *testing.T) {
	var userService = NewService()
	newEmail := "abc@abc.com"
	user := userService.GetAllUsers()[0]
	user.SetEmail(newEmail)
	assert.Equal(t, newEmail, user.email)
}

func TestShouldReturnFirstname(t *testing.T) {
	expectedFirstName := mockUser.firstName
	actualFirstName := mockUser.GetFirstName()

	assert.Equal(t, expectedFirstName, actualFirstName)
}

func TestShouldReturnLastname(t *testing.T) {
	expectedLastName := mockUser.lastName
	actualLastName := mockUser.GetLastName()

	assert.Equal(t, expectedLastName, actualLastName)
}

func TestShouldReturnEmail(t *testing.T) {
	expectedEmail := mockUser.email
	actualEmail := mockUser.GetEmail()

	assert.Equal(t, expectedEmail, actualEmail)
}

func TestShouldReturnPhone(t *testing.T) {
	expectedPhone := mockUser.phone
	actualPhone := mockUser.GetPhone()

	assert.Equal(t, expectedPhone, actualPhone)
}

func TestShouldReturnUsername(t *testing.T) {
	expectedUsername := mockUser.username
	actualUsername := mockUser.GetUsername()

	assert.Equal(t, expectedUsername, actualUsername)
}
