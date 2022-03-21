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

func FindAll() []model.User {

	c := db.Session.DB("shopping_cart_dev").C("user-collection")

	log.Info("Going to find all users")
	var results []model.User
	err := c.Find(nil).All(&results)
	if err != nil {
		log.Error("Error while querying all users ", err)
		panic(err)
	} else {
		log.Info("Found user for this email : ", results)
	}
	return results
}

func SaveUser(user model.User) {

	c := db.Session.DB("shopping_cart_dev").C("user-collection")

	log.Info("Trying to save user : ", user)

	if err := c.Insert(user); err != nil {
		log.Error("Error while saving user ", err)
		panic(err)
	}

	log.Info("User added ")
}

func DeleteUser(email string) bool {

	c := db.Session.DB("shopping_cart_dev").C("user-collection")

	log.Info("Trying to delete user : ", email)

	err := c.Remove(bson.M{"email": email})
	if err != nil {
		log.Error("Error while deleting user with email ", err)
		return false
	} else {
		log.Info("User deleted")
		return true
	}
}

func FindAndUpdate(user model.User) bool {

	c := db.Session.DB("shopping_cart_dev").C("user-collection")

	log.Info("Trying to update user with email : ", user.Email)

	err := c.Update(bson.M{"email": user.Email}, user)
	if err != nil {
		log.Error("Error while updating user ", user, err)
		return false
	} else {
		log.Info("User updated ")
		return true
	}
}
