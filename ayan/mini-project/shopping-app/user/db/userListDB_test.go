package db

import (
	"testing"
	"time"
	"user/domain"

	"github.com/stretchr/testify/assert"
)

func TestNewUserRepositoryListDB(t *testing.T) {

	testUser := NewUser("abc@xyz.com", "aBc&@=+", "Ab Cd", "1, Pqr St.", 951478, "9874563210", "buyer")
	testUser.SetId(1)
	testUser.SetCreatedAt(time.Now())
	testUser.SetUpdatedAt(time.Now())
	testUserList := []User{*testUser}
	testUserRepositoryListDB := &userRepositoryListDB{userList: testUserList}

	newUserRepositoryListDB := NewUserRepositoryListDB(testUserList)

	assert.EqualValues(t, testUserRepositoryListDB, newUserRepositoryListDB)
}

func TestSave(t *testing.T) {

	testUser := *domain.NewUser("abc@xyz.com", "aBc&@=+", "Ab Cd", "1, Pqr St.", 951478, "9874563210", "buyer")

	testUserRepositoryListDB := NewUserRepositoryListDB([]User{})

	savedUser, err := testUserRepositoryListDB.Save(testUser)

	assert.EqualValues(t, testUser, savedUser)
	assert.Nil(t, err)
}

func TestFindUserByEmailForPresentUser(t *testing.T) {

	testDbUser := NewUser("abc@xyz.com", "aBc&@=+", "Ab Cd", "1, Pqr St.", 951478, "9874563210", "buyer")
	testDbUser.SetId(1)
	testDbUser.SetCreatedAt(time.Now())
	testDbUser.SetUpdatedAt(time.Now())
	testUserRepositoryListDB := NewUserRepositoryListDB([]User{*testDbUser})

	testDomainUser := domain.NewUser("abc@xyz.com", "aBc&@=+", "Ab Cd", "1, Pqr St.", 951478, "9874563210", "buyer")

	newDomainUser, err := testUserRepositoryListDB.FindUserByEmail("abc@xyz.com")

	assert.EqualValues(t, testDomainUser, newDomainUser)
	assert.Nil(t, err)
}

func TestFindUserByEmailForNonExistentUser(t *testing.T) {

	testDbUser := NewUser("abc@xyz.com", "aBc&@=+", "Ab Cd", "1, Pqr St.", 951478, "9874563210", "buyer")
	testDbUser.SetId(1)
	testDbUser.SetCreatedAt(time.Now())
	testDbUser.SetUpdatedAt(time.Now())
	testUserRepositoryListDB := NewUserRepositoryListDB([]User{*testDbUser})

	newDomainUser, err := testUserRepositoryListDB.FindUserByEmail("def@xyz.com")

	assert.Nil(t, newDomainUser)
	assert.EqualError(t, err, "user does not exist")
}
