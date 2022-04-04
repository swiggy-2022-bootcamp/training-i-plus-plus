package infra

import (
	"errors"
	"fmt"
	"github.com/swiggy-2022-bootcamp/training-i-plus-plus/murtaza/mini-project/inventory/domain"
	"github.com/swiggy-2022-bootcamp/training-i-plus-plus/murtaza/mini-project/inventory/utils/errs"
	"github.com/swiggy-2022-bootcamp/training-i-plus-plus/murtaza/mini-project/inventory/utils/logger"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2"
	"time"

	"gopkg.in/mgo.v2/bson"
)

const DBUrl = "mongodb://127.0.0.1:27018/items_db"

type itemMongoRepository struct {
	Session *mgo.Session
	Mongo   *mgo.DialInfo
}

func (imr itemMongoRepository) InsertItem(newItem domain.Item) (domain.Item, *errs.AppError) {
	items := imr.Session.DB(imr.Mongo.Database).C(ItemCollectionName)
	mongoItem := imr.toPersistedMongoEntity(newItem)
	if err := items.Insert(mongoItem); err != nil {
		logger.Error("Failed to create Item", zap.Error(err))
		return domain.Item{}, errs.NewUnexpectedError(err.Error())
	}
	return *mongoItem.toDomainEntity(), nil
}

func (imr itemMongoRepository) FindItemById(id int) (*domain.Item, *errs.AppError) {
	items := imr.Session.DB(imr.Mongo.Database).C(ItemCollectionName)
	var result ItemModel
	err := items.Find(bson.M{"id": id}).One(&result)
	if err != nil {
		errMessage := fmt.Sprintf("No item found with itemId: %d", id)
		logger.Error(errMessage, zap.Int("itemId", id), zap.Error(err))
		return nil, errs.NewNotFoundError(err.Error())
	}

	return result.toDomainEntity(), nil
}

func (imr itemMongoRepository) FindItemByName(name string) (*domain.Item, *errs.AppError) {
	items := imr.Session.DB(imr.Mongo.Database).C(ItemCollectionName)
	var result ItemModel
	err := items.Find(bson.M{"name": name}).One(&result)
	if err != nil {
		errMessage := fmt.Sprintf("No item found with name: %d", name)
		logger.Error(errMessage, zap.String("name", name), zap.Error(err))
		return nil, errs.NewNotFoundError(err.Error())
	}

	return result.toDomainEntity(), nil
}

func (imr itemMongoRepository) DeleteItemById(id int) *errs.AppError {
	items := imr.Session.DB(imr.Mongo.Database).C(ItemCollectionName)
	err := items.Remove(bson.M{"id": id})

	if errors.Is(err, mgo.ErrNotFound) {
		errMessage := fmt.Sprintf("No Item found with itemId: %d", id)
		logger.Error(errMessage, zap.Int("itemId", id), zap.Error(err))
		return errs.NewNotFoundError(err.Error())
	}

	if err != nil {
		errMessage := fmt.Sprintf("Cannot delete item with itemId: %d", id)
		logger.Error(errMessage, zap.Int("itemId", id), zap.Error(err))
		return errs.NewUnexpectedError(err.Error())
	}
	return nil
}

func (imr itemMongoRepository) UpdateItem(item domain.Item) (*domain.Item, *errs.AppError) {
	items := imr.Session.DB(imr.Mongo.Database).C(ItemCollectionName)

	change := mgo.Change{
		Update: bson.M{
			"$set": bson.M{
				"name":        item.Name,
				"description": item.Description,
				"quantity":    item.Quantity,
				"updated_at":  time.Now(),
			},
		},
	}

	var updatedItem domain.Item
	_, err := items.Find(bson.M{"id": item.Id}).Apply(change, &updatedItem)

	if err != nil {
		errMessage := fmt.Sprintf("Cannot update item with itemId: %d", item.Id)
		logger.Error(errMessage, zap.Int("itemId", item.Id), zap.Error(err))
		return nil, errs.NewUnexpectedError(err.Error())
	}
	return &item, nil
}

func (imr itemMongoRepository) UpdateItemQuantity(itemId int, quantity int) *errs.AppError {
	items := imr.Session.DB(imr.Mongo.Database).C(ItemCollectionName)

	change := mgo.Change{
		Update: bson.M{
			"$set": bson.M{
				"quantity":   quantity,
				"updated_at": time.Now(),
			},
		},
	}

	var updatedItem domain.Item
	_, err := items.Find(bson.M{"id": itemId}).Apply(change, &updatedItem)

	if err != nil {
		errMessage := fmt.Sprintf("Cannot update item with itemId: %d", itemId)
		logger.Error(errMessage, zap.Int("itemId", itemId), zap.Error(err))
		return errs.NewUnexpectedError(err.Error())
	}
	return nil
}

func NewItemMongoRepository() domain.ItemRepository {

	fmt.Println("Connecting to ", DBUrl)
	mongo, err := mgo.ParseURL(DBUrl)
	s, err := mgo.Dial(DBUrl)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		panic(err.Error())
	}
	s.SetSafe(&mgo.Safe{})
	fmt.Println("Connected to", DBUrl)

	return itemMongoRepository{
		Session: s,
		Mongo:   mongo,
	}
}

func (imr itemMongoRepository) toPersistedMongoEntity(u domain.Item) *ItemModel {
	var nextId = imr.getNextSequence("itemId")
	return &ItemModel{
		Id:          nextId,
		Name:        u.Name,
		Quantity:    u.Quantity,
		Description: u.Description,
		Price:       u.Price,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (imr itemMongoRepository) getNextSequence(seqName string) int {

	type sequenceDoc struct {
		Id            string `bson:"_id"`
		SequenceValue int    `bson:"sequence_value"`
	}

	var seq sequenceDoc
	counters := imr.Session.DB(imr.Mongo.Database).C(CountersCollectionName)

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
