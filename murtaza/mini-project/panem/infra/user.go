package infra

import (
	"time"

	"panem/domain"
)

type UserPersistedEntity struct {
	id        int
	firstName string
	lastName  string
	username  string
	password  string
	phone     string
	email     string
	role      domain.Role
	createdAt time.Time
	updatedAt time.Time
}

// ----- getters and setters --------
func (u UserPersistedEntity) Id() int {
	return u.id
}

func (u UserPersistedEntity) Email() string {
	return u.email
}

func (u UserPersistedEntity) FirstName() string {
	return u.firstName
}

func (u UserPersistedEntity) LastName() string {
	return u.lastName
}

func (u UserPersistedEntity) Username() string {
	return u.username
}

func (u UserPersistedEntity) Password() string {
	return u.password
}

func (u UserPersistedEntity) Phone() string {
	return u.phone
}

func (u UserPersistedEntity) CreatedAt() time.Time {
	return u.createdAt
}

func (u UserPersistedEntity) UpdatedAt() time.Time {
	return u.updatedAt
}

func (u *UserPersistedEntity) SetUpdatedAt(time time.Time) {
	u.updatedAt = time
}

func (u UserPersistedEntity) toDomainEntity() *domain.User {
	domainUser := domain.NewUser(u.firstName, u.lastName, u.username, u.phone, u.email, u.password, u.role)
	domainUser.SetId(u.Id())
	return domainUser
}

func NewUserPersistedEntity(u domain.User) *UserPersistedEntity {
	currentTime := time.Now()
	return &UserPersistedEntity{
		id:        0,
		firstName: u.FirstName,
		lastName:  u.LastName,
		username:  u.Username,
		password:  u.Password,
		email:     u.Email,
		phone:     u.Phone,
		role:      u.Role,
		createdAt: currentTime,
		updatedAt: currentTime,
	}
}
