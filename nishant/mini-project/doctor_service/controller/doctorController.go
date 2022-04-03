package controller

import (
	"fmt"
	"time"

	"github.com/swiggy-2022-bootcamp/training-i-plus-plus/nishant/mini-project/doctor_service/db"
	models "github.com/swiggy-2022-bootcamp/training-i-plus-plus/nishant/mini-project/doctor_service/model"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

// Doctor
// @Description Doctor Request
type DoctorRequest struct {
	Name          string `json:"name" form:"name" bson:"name"`
	Qualification string `json:"qualification" form:"qualification" bson:"qualification"`
}

// Create godoc
// @Summary create doctor
// @Description create doctor
// @Tags Doctor
// @Param   doctor     body DoctorRequest true  "doctor info"
// @Accept  json
// @Success 200
// @Failure 500
// @Router /doctor [post]
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

// Read godoc
// @Summary fetch all  doctor
// @Description fetch all doctor
// @Tags Doctor
// @Accept  json
// @Success 200  {array} models.Doctor{}
// @Failure 500
// @Router /doctor [get]
func Read(c *gin.Context) {
	doctor := db.Session.DB(db.Mongo.Database).C(models.DoctorCollectionName)
	result := []models.Doctor{}
	if err := doctor.Find(nil).All(&result); err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, &result)
}

// Update godoc
// @Summary  Update doctor
// @Description Update doctor
// @Tags Doctor
// @Param   _id path string true "id"
// @Param   doctor      body DoctorRequest true  "doctor info"
// @Accept  json
// @Success 200
// @Failure 500
// @Router /doctor/{_id} [patch]
func Update(c *gin.Context) {

	id := c.Param("_id")
	if id == "" {
		c.Error(fmt.Errorf("id not found"))
	}

	doctor := db.Session.DB(db.Mongo.Database).C(models.DoctorCollectionName)
	toUpdate := DoctorRequest{}
	if err := c.Bind(&toUpdate); err != nil {
		c.Error(err)
		return
	}
	upd := models.Doctor{
		Name:          toUpdate.Name,
		Qualification: toUpdate.Qualification,
		UpdatedOn:     now(),
	}
	if err := doctor.UpdateId(bson.ObjectIdHex(id), upd); err != nil {
		c.Error(err)
		return
	}

	c.Status(200)
}

// Delete godoc
// @Summary Delete doctor by id
// @Tags Doctor
// @Param _id path string true "id"
// @Success 200
// @Failure 500
// @Router /user/{_id} [delete]
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
