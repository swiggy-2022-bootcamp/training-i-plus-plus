package mongodao

import (
	model "auth/internal/dao/mongodao/models"
	"auth/internal/errors"
	"context"
	"sync"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDAO interface {
	AddUser(ctx context.Context, user model.User) *errors.ServerError
}

type mongoDAO struct {
	client *mongo.Client
}

var mongoDao MongoDAO
var mongoDaoOnce sync.Once

func InitMongoDAO(client *mongo.Client) MongoDAO {
	mongoDaoOnce.Do(func() {
		mongoDao = &mongoDAO{
			client: client,
		}
	})

	return mongoDao
}

func GetMongoDAO() MongoDAO {
	if mongoDao == nil {
		panic("Mongo DAO not initialised")
	}

	return mongoDao
}

func (dao *mongoDAO) AddUser(ctx context.Context, user model.User) *errors.ServerError {
	userCollection := dao.client.Database("shopKart").Collection("users")
	ra, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		log.WithError(err).Error("an error occured while inserting a new user in database")
		return &errors.DatabaseInsertionError
	}

	if ra == nil {
		log.WithError(err).Error("user inserted: ", ra, " expected 1")
		return &errors.DatabaseNoInsertionError
	}

	log.Info("user created successfully")
	return nil
}
