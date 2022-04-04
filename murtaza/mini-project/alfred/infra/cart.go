package infra

import (
	"alfred/domain"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	CartCollectionName     = "carts"
	CountersCollectionName = "counters"
)

type Cart struct {
	mongoId   bson.ObjectId  `bson:"_id,omitempty"`
	Id        int            `bson:"id"`
	UserId    int            `bson:"user_id"`
	Items     map[string]int `bson:"items"`
	CreatedAt time.Time      `bson:"created_at"`
	UpdatedAt time.Time      `bson:"updated_at"`
}

func (c Cart) toDomainEntity() *domain.Cart {
	domainCart := domain.NewCart(c.UserId, c.Items)
	domainCart.Id = c.Id
	return domainCart
}
