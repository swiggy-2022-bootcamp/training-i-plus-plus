package db

import (
	"testing"
	"time"

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

func TestSetId(t *testing.T) {

	testUser := User{
		email:    "abc@xyz.com",
		password: "aBc&@=+",
		name:     "Ab Cd",
		address:  "1, Pqr St.",
		zipcode:  951478,
		mobileNo: "9874563210",
		role:     "buyer",
	}

	id := 1
	testUser.SetId(1)

	assert.Equal(t, id, testUser.id)
}

func TestSetCreatedAt(t *testing.T) {

	testUser := User{
		email:    "abc@xyz.com",
		password: "aBc&@=+",
		name:     "Ab Cd",
		address:  "1, Pqr St.",
		zipcode:  951478,
		mobileNo: "9874563210",
		role:     "buyer",
	}

	currTime := time.Now()
	testUser.SetCreatedAt(currTime)

	assert.Equal(t, currTime, testUser.createdAt)
}

func TestSetUpdatedAt(t *testing.T) {

	testUser := User{
		email:    "abc@xyz.com",
		password: "aBc&@=+",
		name:     "Ab Cd",
		address:  "1, Pqr St.",
		zipcode:  951478,
		mobileNo: "9874563210",
		role:     "buyer",
	}

	currTime := time.Now()
	testUser.SetUpdatedAt(currTime)

	assert.Equal(t, currTime, testUser.updatedAt)
}
