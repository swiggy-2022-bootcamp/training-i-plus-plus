package infra

import (
	"alfred/domain"
	"alfred/utils/errs"
	"alfred/utils/logger"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const DBUrl = "mongodb://127.0.0.1:27019/carts_db"

type CartRepository struct {
	Session *mgo.Session
	Mongo   *mgo.DialInfo
}

func NewCartRepository() CartRepository {

	fmt.Println("Connecting to ", DBUrl)
	mongo, err := mgo.ParseURL(DBUrl)
	s, err := mgo.Dial(DBUrl)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Can't connect to mongo, go error %v\n", err), zap.Error(err))
		panic(err.Error())
	}
	s.SetSafe(&mgo.Safe{})
	logger.Info(fmt.Sprintf("Connected to: %s", DBUrl))

	return CartRepository{
		Session: s,
		Mongo:   mongo,
	}
}

func (cr CartRepository) AddToCart(userId int, items map[string]int) *errs.AppError {
	carts := cr.Session.DB(cr.Mongo.Database).C(CartCollectionName)
	var persistedCart Cart
	err := carts.Find(bson.M{"user_id": userId}).One(&persistedCart)
	if errors.Is(err, mgo.ErrNotFound) {
		newCart := domain.Cart{
			UserId: userId,
			Items:  items,
		}

		if err := carts.Insert(cr.toPersistedEntity(newCart)); err != nil {
			logger.Error("Failed to create Cart", zap.Error(err))
			return errs.NewUnexpectedError(err.Error())
		}
	} else if err == nil {
		for itemId := range items {
			if val, ok := persistedCart.Items[itemId]; ok {
				persistedCart.Items[itemId] = val + items[itemId]
			} else {
				persistedCart.Items[itemId] = items[itemId]
			}
		}

		change := mgo.Change{
			Update: bson.M{
				"$set": bson.M{
					"items":     persistedCart.Items,
					"updatedAt": time.Now(),
				},
			},
		}

		var updatedCart domain.Cart
		_, err := carts.Find(bson.M{"id": persistedCart.Id}).Apply(change, &updatedCart)

		if err != nil {
			errMessage := fmt.Sprintf("Cannot update cart with cart Id: %d", persistedCart.Id)
			logger.Error(errMessage, zap.Int("cartId", persistedCart.Id), zap.Error(err))
			return errs.NewUnexpectedError(err.Error())
		}
	}
	return nil
}

func (cr CartRepository) GetCart(userId int) (*domain.Cart, *errs.AppError) {
	carts := cr.Session.DB(cr.Mongo.Database).C(CartCollectionName)
	var persistedCart Cart
	err := carts.Find(bson.M{"user_id": userId}).One(&persistedCart)

	if errors.Is(err, mgo.ErrNotFound) {
		return nil, errs.NewNotFoundError(fmt.Sprintf("No Cart exists for user Id: %d", userId))
	}
	domainCart := persistedCart.toDomainEntity()
	return domainCart, nil
}

func (cr CartRepository) RemoveCart(userId int) *errs.AppError {
	carts := cr.Session.DB(cr.Mongo.Database).C(CartCollectionName)
	err := carts.Remove(bson.M{"user_id": userId})

	if errors.Is(err, mgo.ErrNotFound) {
		return errs.NewNotFoundError(fmt.Sprintf("No Cart exists for user Id: %d", userId))
	}
	return nil
}

func (cr CartRepository) toPersistedEntity(c domain.Cart) *Cart {
	var nextId = cr.getNextSequence("itemId")
	return &Cart{
		Id:        nextId,
		UserId:    c.UserId,
		Items:     c.Items,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (cr CartRepository) getNextSequence(seqName string) int {

	type sequenceDoc struct {
		Id            string `bson:"_id"`
		SequenceValue int    `bson:"sequence_value"`
	}

	var seq sequenceDoc
	counters := cr.Session.DB(cr.Mongo.Database).C(CountersCollectionName)

	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"sequence_value": 1}},
		ReturnNew: true,
		Upsert:    true,
	}

	_, err := counters.Find(bson.M{"_id": seqName}).Apply(change, &seq)
	if err != nil {
		return 0
	} else {
		return seq.SequenceValue
	}
}
