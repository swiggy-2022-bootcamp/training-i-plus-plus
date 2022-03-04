package infra

import (
	"mini-project/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestShouldUpdateUpdatedAt(t *testing.T) {
	firstName := "Murtaza"
	lastName := "Sadriwala"
	phone := "9900887766"
	email := "murtaza896@gmail.com"
	username := "murtaza896"
	password := "Pass!23"
	role := domain.Admin

	user := domain.NewUser(firstName, lastName, username, phone, email, password, role)
	userPersistedEntity := NewUserPersistedEntity(*user)

	newUpdatedAt := time.Now().Add(2)
	userPersistedEntity.SetUpdatedAt(newUpdatedAt)
	assert.Equal(t, newUpdatedAt, userPersistedEntity.UpdatedAt())
}

func TestShouldCreateNewUserPersistedEntity(t *testing.T) {
	firstName := "Murtaza"
	lastName := "Sadriwala"
	phone := "9900887766"
	email := "murtaza896@gmail.com"
	username := "murtaza896"
	password := "Pass!23"
	role := domain.Admin

	user := domain.NewUser(firstName, lastName, username, phone, email, password, role)
	userPersistedEntity := NewUserPersistedEntity(*user)

	assert.Equal(t, user.FirstName(), userPersistedEntity.FirstName())
	assert.Equal(t, user.LastName(), userPersistedEntity.LastName())
	assert.Equal(t, user.Username(), userPersistedEntity.Username())
	assert.Equal(t, user.Email(), userPersistedEntity.Email())
	assert.Equal(t, user.Password(), userPersistedEntity.Password())
	assert.Equal(t, user.Phone(), userPersistedEntity.Phone())
	assert.NotNil(t, userPersistedEntity.CreatedAt())
	assert.NotNil(t, userPersistedEntity.CreatedAt())
	assert.Equal(t, userPersistedEntity.CreatedAt(), userPersistedEntity.UpdatedAt())
}

func TestShouldConvertUserPersistedEntityToDomainEntity(t *testing.T) {
	firstName := "Murtaza"
	lastName := "Sadriwala"
	phone := "9900887766"
	email := "murtaza896@gmail.com"
	username := "murtaza896"
	password := "Pass!23"
	role := domain.Admin

	user := domain.NewUser(firstName, lastName, username, phone, email, password, role)
	userPersistedEntity := NewUserPersistedEntity(*user)

	domainUser := userPersistedEntity.toDomainEntity()

	assert.Equal(t, user.FirstName(), domainUser.FirstName())
	assert.Equal(t, user.LastName(), domainUser.LastName())
	assert.Equal(t, user.Username(), domainUser.Username())
	assert.Equal(t, user.Email(), domainUser.Email())
	assert.Equal(t, user.Password(), domainUser.Password())
	assert.Equal(t, user.Phone(), domainUser.Phone())
}
