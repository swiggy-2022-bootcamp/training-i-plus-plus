package infra

import (
	"panem/domain"
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	// CollectionArticle holds the name of the articles collection
	UserCollectionName     = "users"
	CountersCollectionName = "counters"
)

type UserModel struct {
	mongoId   bson.ObjectId `bson:"_id,omitempty"`
	Id        int
	FirstName string
	LastName  string
	Username  string
	Password  string
	Phone     string
	Email     string
	Role      domain.Role
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u UserModel) toDomainEntity() *domain.User {
	domainUser := domain.NewUser(u.FirstName, u.LastName, u.Username, u.Phone, u.Email, u.Password, u.Role)
	domainUser.Id = u.Id
	return domainUser
}
