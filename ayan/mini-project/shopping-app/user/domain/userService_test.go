package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUserService(t *testing.T) {

	testDefaultUserSvc := &DefaultUserService{
		userDB: nil,
	}

	newDefaultUserSvc := NewUserService(nil)

	assert.EqualValues(t, testDefaultUserSvc, newDefaultUserSvc)
}

func TestRegisterForNewUser(t *testing.T) {

	// testUserRepositoryDB := NewUserRepositoryDB([]db.User{})

}

func TestRegisterForExistingUser(t *testing.T) {

	// testUserRepositoryDB := NewUserRepositoryDB([]db.User{})

}

func TestVerifyCredentialsForGoodCredentials(t *testing.T) {

	// testUserRepositoryDB := NewUserRepositoryDB([]db.User{})

}

func TestVerifyCredentialsForBadCredentials(t *testing.T) {

	// testUserRepositoryDB := NewUserRepositoryDB([]db.User{})

}

func TestVerifyUserCredentialsForInvalidUser(t *testing.T) {

	// testUserRepositoryDB := NewUserRepositoryDB([]db.User{})

}

func TestFindUserByEmailForPresentUser(t *testing.T) {

	// testUserRepositoryDB := NewUserRepositoryDB([]db.User{})

}

func TestFindUserByEmailForNonExistentUser(t *testing.T) {

	// testUserRepositoryDB := NewUserRepositoryDB([]db.User{})

}

func TestVerifyTokenForValidUser(t *testing.T) {

	// testUserRepositoryDB := NewUserRepositoryDB([]db.User{})

}

func TestVerifyTokenForInvalidUser(t *testing.T) {

	// testUserRepositoryDB := NewUserRepositoryDB([]db.User{})

}
