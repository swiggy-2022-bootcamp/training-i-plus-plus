package db

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
	"sample.akash.com/log"
	"sample.akash.com/model"
)

const (
	// MongoDBUrl is the default mongodb url that will be used to connect to the database.
	MongoDBUrl = "mongodb://localhost:2717/shopping_cart_dev"
)

var (
	// Session stores mongo session
	Session *mgo.Session

	// Mongo stores the mongodb connection string information
	Mongo *mgo.DialInfo
)

type repo struct{}

func NewMongoRepository() CustomerRepository {
	return &repo{}
}

func (*repo) Connect() {
	uri := os.Getenv("MONGODB_URL")

	if len(uri) == 0 {
		uri = MongoDBUrl
	}

	mongo, err := mgo.ParseURL(uri)
	s, err := mgo.Dial(uri)
	if err != nil {
		log.Error("Can't connect to mongo, go error %v\n", err)
		panic(err.Error())
	}
	s.SetSafe(&mgo.Safe{})
	log.Info("Connected to", uri)
	Session = s
	Mongo = mongo
}

func (*repo) FindOneWithUsername(username string) *model.User {

	c := Session.DB("shopping_cart_dev").C("user-collection")

	data := &model.User{}
	err := c.Find(bson.M{"username": username}).One(data)
	if err != nil {
		log.Error("Error while finding user with username ", username)
		return nil
	}

	log.Info("Found user for this username : ", *data)

	return data
}

func (*repo) FindAll() []model.User {

	c := Session.DB("shopping_cart_dev").C("user-collection")

	log.Info("Going to find all users")
	var results []model.User
	err := c.Find(nil).All(&results)
	if err != nil {
		log.Error("Error while querying all users ", err)
		panic(err)
	} else {
		log.Info("Found users : ", results)
	}
	return results
}

func (*repo) SaveUser(user model.User) {

	c := Session.DB("shopping_cart_dev").C("user-collection")

	log.Info("Trying to save user : ", user)

	if err := c.Insert(user); err != nil {
		log.Error("Error while saving user ", err)
		panic(err)
	}

	log.Info("User added ")
}

func (*repo) DeleteUser(username string) bool {

	c := Session.DB("shopping_cart_dev").C("user-collection")

	log.Info("Trying to delete user : ", username)

	err := c.Remove(bson.M{"username": username})
	if err != nil {
		log.Error("Error while deleting user with username ", err)
		return false
	} else {
		log.Info("User deleted")
		return true
	}
}

func (*repo) FindAndUpdate(user model.User) bool {

	c := Session.DB("shopping_cart_dev").C("user-collection")

	log.Info("Trying to update user with username : ", user.Username)

	err := c.Update(bson.M{"username": user.Username}, user)
	if err != nil {
		log.Error("Error while updating user ", user, err)
		return false
	} else {
		log.Info("User updated ")
		return true
	}
}
