package infra

import (
	"fmt"
	"panem/domain"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const DBUrl = "mongodb://127.0.0.1:27017/crud_test"

type userMongoRepository struct {
	Session *mgo.Session
	Mongo   *mgo.DialInfo
}

func (umr userMongoRepository) InsertUser(newUser domain.User) (domain.User, error) {
	users := umr.Session.DB(umr.Mongo.Database).C(UserCollectionName)
	mongoUser := umr.toPersistedMongoEntity(newUser)
	if err := users.Insert(mongoUser); err != nil {
		return domain.User{}, fmt.Errorf("could not insert user")
	}
	return *mongoUser.toDomainEntity(), nil
}

func (umr userMongoRepository) FindUserById(id int) (*domain.User, error) {
	users := umr.Session.DB(umr.Mongo.Database).C(UserCollectionName)
	var result UserModel
	err := users.Find(bson.M{"id": id}).One(&result)
	if err != nil {
		return nil, err
	}

	return result.toDomainEntity(), err
}

func (umr userMongoRepository) DeleteUserByUserId(id int) error {
	users := umr.Session.DB(umr.Mongo.Database).C(UserCollectionName)
	err := users.Remove(bson.M{"id": id})
	if err != nil {
		return err
	}
	return nil
}

func (umr userMongoRepository) FindUserByUsername(username string) (*domain.User, error) {
	users := umr.Session.DB(umr.Mongo.Database).C(UserCollectionName)
	var result UserModel
	err := users.Find(bson.M{"username": username}).One(&result)
	if err != nil {
		return nil, err
	}
	return result.toDomainEntity(), err
}

func (umr userMongoRepository) UpdateUser(user domain.User) (*domain.User, error) {
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
		return nil, err
	}
	return &updatedUser, nil
}

func NewUserMongoRepository() userMongoRepository {

	fmt.Println("Connecting to ", DBUrl)
	mongo, err := mgo.ParseURL(DBUrl)
	s, err := mgo.Dial(DBUrl)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		panic(err.Error())
	}
	s.SetSafe(&mgo.Safe{})
	fmt.Println("Connected to", DBUrl)

	return userMongoRepository{
		Session: s,
		Mongo:   mongo,
	}
}

func (umr userMongoRepository) toPersistedMongoEntity(u domain.User) *UserModel {
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

func (umr userMongoRepository) getNextSequence(seqName string) int {

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
