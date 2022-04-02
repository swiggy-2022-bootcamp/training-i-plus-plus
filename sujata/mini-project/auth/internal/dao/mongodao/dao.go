package mongodao

import (
	"auth/config"
	model "auth/internal/dao/mongodao/models"
	"auth/internal/errors"
	"context"
	"sync"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDAO interface {
	AddUser(ctx context.Context, user model.User) *errors.ServerError
	FindUserByEmail(ctx context.Context, email string) (model.User, *errors.ServerError)
}

type mongoDAO struct {
	client *mongo.Client
	config *config.WebServerConfig
}

var mongoDao MongoDAO
var mongoDaoOnce sync.Once

func InitMongoDAO(client *mongo.Client, config *config.WebServerConfig) MongoDAO {
	mongoDaoOnce.Do(func() {
		mongoDao = &mongoDAO{
			client: client,
			config: config,
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
	userCollection := dao.client.Database(dao.config.Db).Collection(dao.config.DbCollection)
	ra, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		log.WithError(err).Error("an error occured while inserting a new user in database")
		return &errors.DatabaseInsertionError
	}

	if ra == nil {
		log.WithError(err).Error("user inserted: 0, expected 1")
		return &errors.DatabaseNoInsertionError
	}

	log.Info("user created successfully")
	return nil
}

func (dao *mongoDAO) FindUserByEmail(ctx context.Context, email string) (model.User, *errors.ServerError) {
	userCollection := dao.client.Database(dao.config.Db).Collection(dao.config.DbCollection)

	userFilter := bson.M{"email": email}
	singleResult := userCollection.FindOne(ctx, userFilter)

	user := model.User{}
	err := singleResult.Decode(&user)
	if err != nil {
		log.WithError(err).Error("error while decoding user into struct from mongodb")
		return user, &errors.UserNotFoundError
	}

	return user, nil
}
