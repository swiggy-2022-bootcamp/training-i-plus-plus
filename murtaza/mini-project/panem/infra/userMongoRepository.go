package infra

import (
	"fmt"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"panem/domain"
	"panem/utils/errs"
	"panem/utils/logger"
	"time"
)

const DBUrl = "mongodb://127.0.0.1:27017/crud_test"

type UserMongoRepository struct {
	Session *mgo.Session
	Mongo   *mgo.DialInfo
}

func (umr UserMongoRepository) InsertUser(newUser domain.User) (domain.User, *errs.AppError) {
	users := umr.Session.DB(umr.Mongo.Database).C(UserCollectionName)
	mongoUser := umr.toPersistedMongoEntity(newUser)
	if err := users.Insert(mongoUser); err != nil {
		logger.Error("Failed to create user", zap.Error(err))
		return domain.User{}, errs.NewUnexpectedError(err.Error())
	}
	return *mongoUser.toDomainEntity(), nil
}

func (umr UserMongoRepository) FindUserById(id int) (*domain.User, *errs.AppError) {
	users := umr.Session.DB(umr.Mongo.Database).C(UserCollectionName)
	var result UserModel
	err := users.Find(bson.M{"id": id}).One(&result)
	if err != nil {
		errMessage := fmt.Sprintf("No user found with userId: %d", id)
		logger.Error(errMessage, zap.Int("userId", id), zap.Error(err))
		return nil, errs.NewNotFoundError(err.Error())
	}

	return result.toDomainEntity(), nil
}

func (umr UserMongoRepository) DeleteUserByUserId(id int) *errs.AppError {
	users := umr.Session.DB(umr.Mongo.Database).C(UserCollectionName)
	err := users.Remove(bson.M{"id": id})
	if err != nil {
		errMessage := fmt.Sprintf("Cannot delete user with userId: %d", id)
		logger.Error(errMessage, zap.Int("userId", id), zap.Error(err))
		return errs.NewUnexpectedError(err.Error())
	}
	return nil
}

func (umr UserMongoRepository) FindUserByUsername(username string) (*domain.User, *errs.AppError) {
	users := umr.Session.DB(umr.Mongo.Database).C(UserCollectionName)
	var result UserModel
	err := users.Find(bson.M{"username": username}).One(&result)
	if err != nil {
		errMessage := fmt.Sprintf("Cannot find user with username: %s", username)
		logger.Error(errMessage, zap.String("username", username), zap.Error(err))
		return nil, errs.NewNotFoundError(err.Error())
	}
	return result.toDomainEntity(), nil
}

func (umr UserMongoRepository) UpdateUser(user domain.User) (*domain.User, *errs.AppError) {
	users := umr.Session.DB(umr.Mongo.Database).C(UserCollectionName)

	change := mgo.Change{
		Update: bson.M{
			"$set": bson.M{
				"firstname": user.FirstName,
				"lastname":  user.LastName,
				"username":  user.Username,
				"password":  user.Password,
				"phone":     user.Phone,
				"email":     user.Email,
				"role":      0,
				"updatedat": time.Now(),
			},
		},
	}

	var updatedUser domain.User
	_, err := users.Find(bson.M{"id": user.Id}).Apply(change, &updatedUser)

	if err != nil {
		errMessage := fmt.Sprintf("Cannot update user with UserId: %d", user.Id)
		logger.Error(errMessage, zap.Int("userId", user.Id), zap.Error(err))
		return nil, errs.NewUnexpectedError(err.Error())
	}
	return &updatedUser, nil
}

func NewUserMongoRepository() UserMongoRepository {

	fmt.Println("Connecting to ", DBUrl)
	mongo, err := mgo.ParseURL(DBUrl)
	s, err := mgo.Dial(DBUrl)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Can't connect to mongo, go error %v\n", err), zap.Error(err))
		panic(err.Error())
	}
	s.SetSafe(&mgo.Safe{})
	logger.Info(fmt.Sprintf("Connected to: %s", DBUrl))

	return UserMongoRepository{
		Session: s,
		Mongo:   mongo,
	}
}

func (umr UserMongoRepository) toPersistedMongoEntity(u domain.User) *UserModel {
	var nextId = umr.getNextSequence("userId")
	return &UserModel{
		Id:        nextId,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Phone:     u.Phone,
		Email:     u.Email,
		Username:  u.Username,
		Password:  u.Password,
		Role:      u.Role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (umr UserMongoRepository) getNextSequence(seqName string) int {

	type sequenceDoc struct {
		Id            string `bson:"_id"`
		SequenceValue int    `bson:"sequence_value"`
	}

	var seq sequenceDoc
	counters := umr.Session.DB(umr.Mongo.Database).C(CountersCollectionName)

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
