package user

import (
	"gopkg.in/mgo.v2/bson"
	"sample.akash.com/db"
	"sample.akash.com/log"
	"sample.akash.com/model"
)

func FindOneWithEmail(email string) *model.User {

	c := db.Session.DB("shopping_cart_dev").C("user-collection")

	data := &model.User{}
	err := c.Find(bson.M{"email": email}).One(data)
	if err != nil {
		log.Error("Error while finding user with email ", email)
		return nil
	}

	log.Info("Found user for this email : ", *data)

	return data
}

func SaveUser(user model.User) {

	c := db.Session.DB("shopping_cart_dev").C("user-collection")

	log.Info("Trying to save user : ", user)
	log.Info("add ", &user)

	if err := c.Insert(user); err != nil {
		log.Error("Error while saving user ", err)
		panic(err)
	}

	log.Info("User added ")
}
