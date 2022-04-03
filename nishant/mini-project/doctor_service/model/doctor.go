package models

import "gopkg.in/mgo.v2/bson"

const (
	// CollectionArticle holds the name of the articles collection
	DoctorCollectionName = "doctors"
)

// Doctor
// @Description Doctor
type Doctor struct {
	Id            bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name          string        `json:"name" form:"name" binding:"required" bson:"name"`
	Qualification string        `json:"qualification" form:"qualification" binding:"required" bson:"qualification"`
	UpdatedOn     int64         `json:"updated_on" bson:"updated_on"`
}
