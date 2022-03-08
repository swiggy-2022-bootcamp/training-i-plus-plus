package handler

import (
	"fmt"
	"time"
	"usecase/crud_mongo/db"
	models "usecase/crud_mongo/model"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func Create(c *gin.Context) {

	doctor := db.Session.DB(db.Mongo.Database).C(models.DoctorCollectionName)
	newDoctor := models.Doctor{}
	if err := c.Bind(&newDoctor); err != nil {
		c.Error(err)
		return
	}
	newDoctor.UpdatedOn = now()

	if err := doctor.Insert(newDoctor); err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, &newDoctor)
}

func Read(c *gin.Context) {
	doctor := db.Session.DB(db.Mongo.Database).C(models.DoctorCollectionName)
	result := []models.Doctor{}
	if err := doctor.Find(nil).All(&result); err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, &result)
}

func Update(c *gin.Context) {

	id := c.Param("_id")
	if id == "" {
		c.Error(fmt.Errorf("id not found"))
	}

	doctor := db.Session.DB(db.Mongo.Database).C(models.DoctorCollectionName)
	toUpdate := models.Doctor{}
	if err := c.Bind(&toUpdate); err != nil {
		c.Error(err)
		return
	}
	toUpdate.UpdatedOn = now()
	if err := doctor.UpdateId(bson.ObjectIdHex(id), toUpdate); err != nil {
		c.Error(err)
		return
	}

	c.Status(200)
}

func Delete(c *gin.Context) {

	id := c.Param("_id")
	if id == "" {
		c.Error(fmt.Errorf("id not found"))

	}
	doctor := db.Session.DB(db.Mongo.Database).C(models.DoctorCollectionName)

	if err := doctor.RemoveId(bson.ObjectIdHex(id)); err != nil {
		c.Error(err)
		return
	}

	c.Status(200)
}

func now() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
