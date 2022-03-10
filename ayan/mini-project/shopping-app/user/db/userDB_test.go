package db

import (
	"shopping-app/user/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewUserRepositoryDB(t *testing.T) {

	testUser := NewUser("abc@xyz.com", "aBc&@=+", "Ab Cd", "1, Pqr St.", 951478, "9874563210", "buyer")
	testUser.SetId(1)
	testUser.SetCreatedAt(time.Now())
	testUser.SetUpdatedAt(time.Now())
	testUserList := []User{*testUser}
	testUserRepositoryDB := &userRepositoryDB{userList: testUserList}

	newUserRepositoryDB := NewUserRepositoryDB(testUserList)

	assert.EqualValues(t, testUserRepositoryDB, newUserRepositoryDB)
}

func TestSave(t *testing.T) {

	testUser := *domain.NewUser("abc@xyz.com", "aBc&@=+", "Ab Cd", "1, Pqr St.", 951478, "9874563210", "buyer")

	testUserRepositoryDB := NewUserRepositoryDB([]User{})

	savedUser, err := testUserRepositoryDB.Save(testUser)

	assert.EqualValues(t, testUser, savedUser)
	assert.Nil(t, err)
}

func TestFindUserByEmailForPresentUser(t *testing.T) {

	testDbUser := NewUser("abc@xyz.com", "aBc&@=+", "Ab Cd", "1, Pqr St.", 951478, "9874563210", "buyer")
	testDbUser.SetId(1)
	testDbUser.SetCreatedAt(time.Now())
	testDbUser.SetUpdatedAt(time.Now())
	testUserRepositoryDB := NewUserRepositoryDB([]User{*testDbUser})

	testDomainUser := domain.NewUser("abc@xyz.com", "aBc&@=+", "Ab Cd", "1, Pqr St.", 951478, "9874563210", "buyer")

	newDomainUser, err := testUserRepositoryDB.FindUserByEmail("abc@xyz.com")

	assert.EqualValues(t, testDomainUser, newDomainUser)
	assert.Nil(t, err)
}

func TestFindUserByEmailForNonExistentUser(t *testing.T) {

	testDbUser := NewUser("abc@xyz.com", "aBc&@=+", "Ab Cd", "1, Pqr St.", 951478, "9874563210", "buyer")
	testDbUser.SetId(1)
	testDbUser.SetCreatedAt(time.Now())
	testDbUser.SetUpdatedAt(time.Now())
	testUserRepositoryDB := NewUserRepositoryDB([]User{*testDbUser})

	newDomainUser, err := testUserRepositoryDB.FindUserByEmail("def@xyz.com")

	assert.Nil(t, newDomainUser)
	assert.EqualError(t, err, "user does not exist")
}
