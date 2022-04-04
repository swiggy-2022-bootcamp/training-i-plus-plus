package infra

import (
	"github.com/swiggy-2022-bootcamp/training-i-plus-plus/murtaza/mini-project/inventory/domain"
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	ItemCollectionName     = "items"
	CountersCollectionName = "counters"
)

type ItemModel struct {
	mongoId     bson.ObjectId `bson:"_id,omitempty"`
	Id          int           `bson:"id"`
	Name        string        `bson:"name"`
	Description string        `bson:"description"`
	Quantity    int           `bson:"quantity"`
	CreatedAt   time.Time     `bson:"created_at"`
	UpdatedAt   time.Time     `bson:"updated_at"`
}

func (u ItemModel) toDomainEntity() *domain.Item {
	domainItem := domain.NewItem(u.Name, u.Description, u.Quantity)
	domainItem.Id = u.Id
	return domainItem
}
