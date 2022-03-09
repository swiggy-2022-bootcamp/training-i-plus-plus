package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {

	testUser := &User{
		email:    "abc@xyz.com",
		password: "aBc&@=+",
		name:     "Ab Cd",
		address:  "1, Pqr St.",
		zipcode:  951478,
		mobileNo: "9874563210",
		role:     "buyer",
	}

	newUser := NewUser("abc@xyz.com", "aBc&@=+", "Ab Cd", "1, Pqr St.", 951478, "9874563210", "buyer")

	assert.EqualValues(t, testUser, newUser)
}

func TestEmail(t *testing.T) {

	testUser := NewUser("abc@xyz.com", "aBc&@=+", "Ab Cd", "1, Pqr St.", 951478, "9874563210", "buyer")

	assert.EqualValues(t, "abc@xyz.com", testUser.Email())
}

func TestPassword(t *testing.T) {

	testUser := NewUser("abc@xyz.com", "aBc&@=+", "Ab Cd", "1, Pqr St.", 951478, "9874563210", "buyer")

	assert.EqualValues(t, "aBc&@=+", testUser.Password())
}

func TestName(t *testing.T) {

	testUser := NewUser("abc@xyz.com", "aBc&@=+", "Ab Cd", "1, Pqr St.", 951478, "9874563210", "buyer")

	assert.EqualValues(t, "Ab Cd", testUser.Name())
}

func TestAddress(t *testing.T) {

	testUser := NewUser("abc@xyz.com", "aBc&@=+", "Ab Cd", "1, Pqr St.", 951478, "9874563210", "buyer")

	assert.EqualValues(t, "1, Pqr St.", testUser.Address())
}

func TestZipcode(t *testing.T) {

	testUser := NewUser("abc@xyz.com", "aBc&@=+", "Ab Cd", "1, Pqr St.", 951478, "9874563210", "buyer")

	assert.EqualValues(t, 951478, testUser.Zipcode())
}

func TestMobileNo(t *testing.T) {

	testUser := NewUser("abc@xyz.com", "aBc&@=+", "Ab Cd", "1, Pqr St.", 951478, "9874563210", "buyer")

	assert.EqualValues(t, "9874563210", testUser.MobileNo())
}

func TestRole(t *testing.T) {

	testUser := NewUser("abc@xyz.com", "aBc&@=+", "Ab Cd", "1, Pqr St.", 951478, "9874563210", "buyer")

	assert.EqualValues(t, "buyer", testUser.Role())
}

func TestSetEmail(t *testing.T) {

	testUser := NewUser("abc@xyz.com", "aBc&@=+", "Ab Cd", "1, Pqr St.", 951478, "9874563210", "buyer")

	testUser.SetEmail("def@xyz.com")

	assert.EqualValues(t, "def@xyz.com", testUser.Email())
}

func TestSetPassword(t *testing.T) {

	testUser := NewUser("abc@xyz.com", "aBc&@=+", "Ab Cd", "1, Pqr St.", 951478, "9874563210", "buyer")

	testUser.SetPassword("564512465&@=+")

	assert.EqualValues(t, "564512465&@=+", testUser.Password())
}

func TestSetName(t *testing.T) {

	testUser := NewUser("abc@xyz.com", "aBc&@=+", "Ab Cd", "1, Pqr St.", 951478, "9874563210", "buyer")

	testUser.SetName("Ba Dc Ef")

	assert.EqualValues(t, "Ba Dc Ef", testUser.Name())
}

func TestSetAddress(t *testing.T) {

	testUser := NewUser("abc@xyz.com", "aBc&@=+", "Ab Cd", "1, Pqr St.", 951478, "9874563210", "buyer")

	testUser.SetAddress("1, ZYX St.")

	assert.EqualValues(t, "1, ZYX St.", testUser.Address())
}

func TestSetZipcode(t *testing.T) {

	testUser := NewUser("abc@xyz.com", "aBc&@=+", "Ab Cd", "1, Pqr St.", 951478, "9874563210", "buyer")

	testUser.SetZipcode(654213)

	assert.EqualValues(t, 654213, testUser.Zipcode())
}

func TestSetMobileNo(t *testing.T) {

	testUser := NewUser("abc@xyz.com", "aBc&@=+", "Ab Cd", "1, Pqr St.", 951478, "9874563210", "buyer")

	testUser.SetMobileNo("0123654789")

	assert.EqualValues(t, "0123654789", testUser.MobileNo())
}

func TestSetRole(t *testing.T) {

	testUser := NewUser("abc@xyz.com", "aBc&@=+", "Ab Cd", "1, Pqr St.", 951478, "9874563210", "buyer")

	testUser.SetRole("seller")

	assert.EqualValues(t, "seller", testUser.Role())
}
