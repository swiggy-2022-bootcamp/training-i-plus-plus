package models

import "gopkg.in/mgo.v2/bson"

// const (
// 	// CollectionArticle holds the name of the articles collection
// 	CollectionArticle = "articles"
// )

// Article model
type User struct {
	Id        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string        `json:"name" form:"name" binding:"required" bson:"name"`
	// Body      string        `json:"body" form:"body" binding:"required" bson:"body"`
	// CreatedOn int64         `json:"created_on" bson:"created_on"`
	// UpdatedOn int64         `json:"updated_on" bson:"updated_on"`
	// User      bson.ObjectId `json:"user"`
}