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
	mongoId         bson.ObjectId      `bson:"_id,omitempty"`
	Id              int                `bson:"id"`
	FirstName       string             `bson:"first_name"`
	LastName        string             `bson:"last_name"`
	Username        string             `bson:"username"`
	Password        string             `bson:"password"`
	Phone           string             `bson:"phone"`
	Email           string             `bson:"email"`
	Role            domain.Role        `bson:"role"`
	PurchaseHistory map[string]float64 `bson:"purchase_history"`
	CreatedAt       time.Time          `bson:"created_at"`
	UpdatedAt       time.Time          `bson:"updated_at"`
}

func (u UserModel) toDomainEntity() *domain.User {
	domainUser := domain.NewUser(u.FirstName, u.LastName, u.Username, u.Phone, u.Email, u.Password, u.Role)
	domainUser.Id = u.Id
	domainUser.PurchaseHistory = u.PurchaseHistory
	return domainUser
}
