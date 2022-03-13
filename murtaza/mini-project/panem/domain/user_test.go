package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldGetRoleString(t *testing.T) {
	role := Admin
	var expected string = "admin"
	var actual string = role.String()

	assert.Equal(t, expected, actual)
}

func TestShouldReturnEnumIndexForRole(t *testing.T) {
	role := Admin
	var expected int = 0
	var actual int = role.EnumIndex()

	assert.Equal(t, expected, actual)
}

func TestShouldGetAdminEnumByIndex(t *testing.T) {
	var expected Role = Admin
	actual, err := GetEnumByIndex(0)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestShouldGetSellerEnumByIndex(t *testing.T) {
	var expected Role = Seller
	actual, err := GetEnumByIndex(1)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestShouldGetCustomerEnumByIndex(t *testing.T) {
	var expected Role = Customer
	actual, err := GetEnumByIndex(2)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestShouldReturnErrOnGetEnumByIndexForInvalidIndex(t *testing.T) {
	var expected Role = -1
	actual, err := GetEnumByIndex(1000)

	assert.Error(t, err)
	assert.Equal(t, expected, actual)
}

func TestShouldReturnNewUser(t *testing.T) {
	firstName := "Murtaza"
	lastName := "Sadriwala"
	phone := "9900887766"
	email := "murtaza896@gmail.com"
	username := "murtaza896"
	password := "Pass!23"
	role := Admin

	user := NewUser(firstName, lastName, username, phone, email, password, role)
	assert.Equal(t, firstName, user.FirstName())
	assert.Equal(t, lastName, user.LastName())
	assert.Equal(t, username, user.Username())
	assert.Equal(t, phone, user.Phone())
	assert.Equal(t, email, user.Email())
	assert.Equal(t, password, user.Password())
}

func TestShouldUpdateEmail(t *testing.T) {
	firstName := "Murtaza"
	lastName := "Sadriwala"
	phone := "9900887766"
	email := "murtaza896@gmail.com"
	username := "murtaza896"
	password := "Pass!23"
	role := Admin

	newEmail := "msadriwala.1198@gmail.com"
	user := NewUser(firstName, lastName, username, phone, email, password, role)

	user.SetEmail(newEmail)

	assert.Equal(t, newEmail, user.Email())
}

func TestShouldUpdatePhone(t *testing.T) {
	firstName := "Murtaza"
	lastName := "Sadriwala"
	phone := "9900887766"
	email := "murtaza896@gmail.com"
	username := "murtaza896"
	password := "Pass!23"
	role := Admin

	newPhone := "9999955555"
	user := NewUser(firstName, lastName, username, phone, email, password, role)

	user.SetPhone(newPhone)

	assert.Equal(t, newPhone, user.Phone())
}

func TestShouldUpdateUsername(t *testing.T) {
	firstName := "Murtaza"
	lastName := "Sadriwala"
	phone := "9900887766"
	email := "murtaza896@gmail.com"
	username := "murtaza896"
	password := "Pass!23"
	role := Admin

	newUsername := "newUsername"
	user := NewUser(firstName, lastName, username, phone, email, password, role)

	user.SetUsername(newUsername)

	assert.Equal(t, newUsername, user.Username())
}

func TestShouldUpdatePassword(t *testing.T) {
	firstName := "Murtaza"
	lastName := "Sadriwala"
	phone := "9900887766"
	email := "murtaza896@gmail.com"
	username := "murtaza896"
	password := "Pass!23"
	role := Admin

	newPassword := "newPass"
	user := NewUser(firstName, lastName, username, phone, email, password, role)

	user.SetPassword(newPassword)

	assert.Equal(t, newPassword, user.Password())
}

func TestShouldUpdateFirstName(t *testing.T) {
	firstName := "Murtaza"
	lastName := "Sadriwala"
	phone := "9900887766"
	email := "murtaza896@gmail.com"
	username := "murtaza896"
	password := "Pass!23"
	role := Admin

	newFirstName := "MurtazaNew"
	user := NewUser(firstName, lastName, username, phone, email, password, role)

	user.SetFirstName(newFirstName)

	assert.Equal(t, newFirstName, user.FirstName())
}

func TestShouldUpdateLastName(t *testing.T) {
	firstName := "Murtaza"
	lastName := "Sadriwala"
	phone := "9900887766"
	email := "murtaza896@gmail.com"
	username := "murtaza896"
	password := "Pass!23"
	role := Admin

	newLastName := "newLastName"
	user := NewUser(firstName, lastName, username, phone, email, password, role)

	user.SetLastName(newLastName)

	assert.Equal(t, newLastName, user.LastName())
}

func TestShouldUpdateRole(t *testing.T) {
	firstName := "Murtaza"
	lastName := "Sadriwala"
	phone := "9900887766"
	email := "murtaza896@gmail.com"
	username := "murtaza896"
	password := "Pass!23"
	role := Admin

	newRole := Seller
	user := NewUser(firstName, lastName, username, phone, email, password, role)

	user.SetRole(newRole)

	assert.Equal(t, newRole, user.Role())
}
