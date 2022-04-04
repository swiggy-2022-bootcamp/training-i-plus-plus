package infra

import (
	"alfred/domain"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	OrderCollectionName = "orders"
)

type Order struct {
	mongoId     bson.ObjectId  `bson:"_id,omitempty"`
	Id          int            `bson:"order_id"`
	UserId      int            `bson:"user_id"`
	OrderAmount float64        `bson:"order_amount"`
	Items       map[string]int `bson:"items"`
	CreatedAt   time.Time      `bson:"created_at"`
	UpdatedAt   time.Time      `bson:"updated_at"`
}

func (o Order) toDomainEntity() *domain.Order {
	order := domain.NewOrder(o.UserId, o.OrderAmount, o.Items)
	order.Id = o.Id
	return order
}
