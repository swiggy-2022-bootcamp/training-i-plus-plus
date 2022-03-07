package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindUserByEmailForPresentUser(t *testing.T) {

	testUser := User{
		email:    "abc@xyz.com",
		password: "aBc&@=+",
		name:     "Ab Cd",
		address:  "1, Pqr St.",
		pincode:  951478,
		mobileNo: 9874563210,
	}

	svc := Service{
		userList: []User{
			testUser,
		},
	}

	user, isUserPresent := svc.FindUserByEmail("abc@xyz.com")

	assert.EqualValues(t, testUser, user)
	assert.True(t, isUserPresent)
}

func TestFindUserByEmailForAbsentUser(t *testing.T) {

	testUser := User{
		email:    "abc@xyz.com",
		password: "aBc&@=+",
		name:     "Ab Cd",
		address:  "1, Pqr St.",
		pincode:  951478,
		mobileNo: 9874563210,
	}

	svc := Service{
		userList: []User{
			testUser,
		},
	}

	blankUser := User{}

	user, isUserPresent := svc.FindUserByEmail("def@xyz.com")

	assert.EqualValues(t, blankUser, user)
	assert.False(t, isUserPresent)
}

func TestAddUser(t *testing.T) {

	testUser := User{
		email:    "abc@xyz.com",
		password: "aBc&@=+",
		name:     "Ab Cd",
		address:  "1, Pqr St.",
		pincode:  951478,
		mobileNo: 9874563210,
	}
	testUserList := []User{testUser}

	svc := Service{}

	svc.AddUser(testUser)

	assert.ElementsMatch(t, testUserList, svc.userList)
}

func TestRegisterUserForNewUser(t *testing.T) {

	testUser := User{
		email:    "abc@xyz.com",
		password: "aBc&@=+",
		name:     "Ab Cd",
		address:  "1, Pqr St.",
		pincode:  951478,
		mobileNo: 9874563210,
	}

	svc := Service{}

	user, err := svc.RegisterUser(testUser)

	assert.EqualValues(t, testUser, *user)
	assert.Nil(t, err)
}

func TestRegisterUserForExistingUser(t *testing.T) {

	testUser := User{
		email:    "abc@xyz.com",
		password: "aBc&@=+",
		name:     "Ab Cd",
		address:  "1, Pqr St.",
		pincode:  951478,
		mobileNo: 9874563210,
	}

	svc := Service{
		userList: []User{
			testUser,
		},
	}

	newUser := User{
		email:    "abc@xyz.com",
		password: "aBc&@=+",
		name:     "Ab Cd",
		address:  "1, Pqr St.",
		pincode:  951478,
		mobileNo: 9874563210,
	}

	user, err := svc.RegisterUser(newUser)

	assert.Nil(t, user)
	assert.EqualError(t, err, "user already exists")
}

func TestVerifyUserCredentialsForGoodCredentials(t *testing.T) {

	testUser := User{
		email:    "abc@xyz.com",
		password: "aBc&@=+",
		name:     "Ab Cd",
		address:  "1, Pqr St.",
		pincode:  951478,
		mobileNo: 9874563210,
	}

	svc := Service{
		userList: []User{
			testUser,
		},
	}

	isValid := svc.VerifyUserCredentials("abc@xyz.com", "aBc&@=+")

	assert.True(t, isValid)
}

func TestVerifyUserCredentialsForBadCredentials(t *testing.T) {

	testUser := User{
		email:    "abc@xyz.com",
		password: "aBc&@=+",
		name:     "Ab Cd",
		address:  "1, Pqr St.",
		pincode:  951478,
		mobileNo: 9874563210,
	}

	svc := Service{
		userList: []User{
			testUser,
		},
	}

	isValid := svc.VerifyUserCredentials("abc@xyz.com", "sfbfnthy")

	assert.False(t, isValid)
}

func TestVerifyUserCredentialsForInvalidUser(t *testing.T) {

	testUser := User{
		email:    "abc@xyz.com",
		password: "aBc&@=+",
		name:     "Ab Cd",
		address:  "1, Pqr St.",
		pincode:  951478,
		mobileNo: 9874563210,
	}

	svc := Service{
		userList: []User{
			testUser,
		},
	}

	isValid := svc.VerifyUserCredentials("def@xyz.com", "aBc&@=+")

	assert.False(t, isValid)
}

func TestLoginUserForValidUser(t *testing.T) {

	testUser := User{
		email:    "abc@xyz.com",
		password: "aBc&@=+",
		name:     "Ab Cd",
		address:  "1, Pqr St.",
		pincode:  951478,
		mobileNo: 9874563210,
	}

	svc := Service{
		userList: []User{
			testUser,
		},
		tokenMap: map[string]string{},
	}

	expectedToken := "$" + testUser.email + "$" + testUser.password + "$"

	token, err := svc.LoginUser("abc@xyz.com", "aBc&@=+")

	assert.Equal(t, expectedToken, token)
	assert.Nil(t, err)
}

func TestLoginUserForIncorrectUser(t *testing.T) {

	testUser := User{
		email:    "abc@xyz.com",
		password: "aBc&@=+",
		name:     "Ab Cd",
		address:  "1, Pqr St.",
		pincode:  951478,
		mobileNo: 9874563210,
	}

	svc := Service{
		userList: []User{
			testUser,
		},
		tokenMap: map[string]string{},
	}

	expectedToken := ""

	token, err := svc.LoginUser("abc@xyz.com", "ert764564")

	assert.Equal(t, expectedToken, token)
	assert.EqualError(t, err, "invalid user credentials")
}

func TestLoginUserForInvalidUser(t *testing.T) {

	testUser := User{
		email:    "abc@xyz.com",
		password: "aBc&@=+",
		name:     "Ab Cd",
		address:  "1, Pqr St.",
		pincode:  951478,
		mobileNo: 9874563210,
	}

	svc := Service{
		userList: []User{
			testUser,
		},
		tokenMap: map[string]string{},
	}

	expectedToken := ""

	token, err := svc.LoginUser("def@xyz.com", "aBc&@=+")

	assert.Equal(t, expectedToken, token)
	assert.EqualError(t, err, "invalid user credentials")
}

func TestVerifySavedUserTokenForLoggedInUser(t *testing.T) {

	testUser := User{
		email:    "abc@xyz.com",
		password: "aBc&@=+",
		name:     "Ab Cd",
		address:  "1, Pqr St.",
		pincode:  951478,
		mobileNo: 9874563210,
	}

	svc := Service{
		userList: []User{
			testUser,
		},
		tokenMap: map[string]string{},
	}

	token, _ := svc.LoginUser("abc@xyz.com", "aBc&@=+")

	isValid := svc.VerifySavedUserToken("abc@xyz.com", token)

	assert.True(t, isValid)
}

func TestVerifySavedUserTokenForIncorrectToken(t *testing.T) {

	testUser := User{
		email:    "abc@xyz.com",
		password: "aBc&@=+",
		name:     "Ab Cd",
		address:  "1, Pqr St.",
		pincode:  951478,
		mobileNo: 9874563210,
	}

	svc := Service{
		userList: []User{
			testUser,
		},
		tokenMap: map[string]string{},
	}

	svc.LoginUser("abc@xyz.com", "aBc&@=+")

	isValid := svc.VerifySavedUserToken("abc@xyz.com", "DummyToken")

	assert.False(t, isValid)
}

func TestVerifySavedUserTokenForLoggedOutUser(t *testing.T) {

	testUser := User{
		email:    "abc@xyz.com",
		password: "aBc&@=+",
		name:     "Ab Cd",
		address:  "1, Pqr St.",
		pincode:  951478,
		mobileNo: 9874563210,
	}

	svc := Service{
		userList: []User{
			testUser,
		},
		tokenMap: map[string]string{},
	}

	token, _ := svc.LoginUser("abc@xyz.com", "aBc&@=+")
	svc.LogoutUser("abc@xyz.com")

	isValid := svc.VerifySavedUserToken("abc@xyz.com", token)

	assert.False(t, isValid)
}

func TestLogoutUserForLoggedInUser(t *testing.T) {

	testUser := User{
		email:    "abc@xyz.com",
		password: "aBc&@=+",
		name:     "Ab Cd",
		address:  "1, Pqr St.",
		pincode:  951478,
		mobileNo: 9874563210,
	}

	svc := Service{
		userList: []User{
			testUser,
		},
		tokenMap: map[string]string{},
	}

	svc.LoginUser("abc@xyz.com", "aBc&@=+")

	isValid, err := svc.LogoutUser("abc@xyz.com")

	assert.True(t, isValid)
	assert.Nil(t, err)
}

func TestLogoutUserForLoggedOutUser(t *testing.T) {

	testUser := User{
		email:    "abc@xyz.com",
		password: "aBc&@=+",
		name:     "Ab Cd",
		address:  "1, Pqr St.",
		pincode:  951478,
		mobileNo: 9874563210,
	}

	svc := Service{
		userList: []User{
			testUser,
		},
		tokenMap: map[string]string{},
	}

	svc.LoginUser("abc@xyz.com", "aBc&@=+")
	svc.LogoutUser("abc@xyz.com")

	isValid, err := svc.LogoutUser("abc@xyz.com")

	assert.False(t, isValid)
	assert.EqualError(t, err, "user not logged in")
}

func TestLogoutUserForInvalidUser(t *testing.T) {

	testUser := User{
		email:    "abc@xyz.com",
		password: "aBc&@=+",
		name:     "Ab Cd",
		address:  "1, Pqr St.",
		pincode:  951478,
		mobileNo: 9874563210,
	}

	svc := Service{
		userList: []User{
			testUser,
		},
		tokenMap: map[string]string{},
	}

	isValid, err := svc.LogoutUser("def@xyz.com")

	assert.False(t, isValid)
	assert.EqualError(t, err, "invalid user email")
}
