package infra

import (
	"fmt"
	"time"

	"github.com/swiggy-2022-bootcamp/training-i-plus-plus/murtaza/mini-project/panem/domain"
)

type userRepository struct {
	db []*UserPersistedEntity
}

func (urdb userRepository) GetAllUsers() ([]*domain.User, error) {
	res := []*domain.User{}
	for _, u := range urdb.db {
		res = append(res, u.toDomainEntity())
	}
	return res, nil
}

func (urdb userRepository) FindByUsername(username string) (*domain.User, error) {
	for _, val := range urdb.db {
		if username == val.Username() {
			return val.toDomainEntity(), nil
		}
	}
	return nil, fmt.Errorf("no users found")
}

func (urdb userRepository) FindByEmail(email string) (*domain.User, error) {
	for _, val := range urdb.db {
		if email == val.Email() {
			return val.toDomainEntity(), nil
		}
	}
	return nil, fmt.Errorf("no users found")
}

func (urdb *userRepository) Save(u domain.User) (domain.User, error) {
	userPersistedEntity := toPersistedEntity(u)
	urdb.db = append(urdb.db, userPersistedEntity)
	return u, nil
}

func (urdb *userRepository) DeleteUserByUsername(username string) (bool, error) {
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

func NewUserRepository() domain.UserRepository {
	return &userRepository{
		db: []*UserPersistedEntity{},
	}
}
