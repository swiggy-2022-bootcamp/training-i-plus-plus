package infra

import (
	"fmt"
	"mini-project/domain"
	"time"
)

type UserRepository struct {
	db []*UserPersistedEntity
}

func (urdb UserRepository) GetAllUsers() []*domain.User {
	res := []*domain.User{}
	for _, u := range urdb.db {
		res = append(res, u.toDomainEntity())
	}
	return res
}

func (urdb UserRepository) FindByUsername(username string) (*domain.User, error) {
	for _, val := range urdb.GetAllUsers() {
		if username == val.Username() {
			return val, nil
		}
	}
	return nil, fmt.Errorf("no users found")
}

func (urdb UserRepository) FindByEmail(email string) (*domain.User, error) {
	for _, val := range urdb.GetAllUsers() {
		if email == val.Email() {
			return val, nil
		}
	}
	return nil, fmt.Errorf("no users found")
}

func (urdb *UserRepository) Save(u domain.User) (domain.User, error) {
	userPersistedEntity := toPersistedEntity(u)
	urdb.db = append(urdb.db, userPersistedEntity)
	return u, nil
}

func (urdb *UserRepository) DeleteUserByUsername(username string) (bool, error) {
	var pos int = -1
	users := urdb.db
	for i, v := range users {
		if v.Username() == username {
			pos = i
			break
		}
	}
	urdb.db = append(users[0:pos], users[pos+1:]...)
	return true, nil
}

func toPersistedEntity(u domain.User) *UserPersistedEntity {
	return &UserPersistedEntity{
		id:        1,
		firstName: u.FirstName(),
		lastName:  u.LastName(),
		phone:     u.Phone(),
		email:     u.Email(),
		username:  u.Username(),
		password:  u.Password(),
		role:      u.Role(),
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
}
