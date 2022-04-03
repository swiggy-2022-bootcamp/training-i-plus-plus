package controller

import (
	"log"
	"time"

	entity "github.com/swiggy-2022-bootcamp/training-i-plus-plus/nishant/mini-project/doctor_service/api_entities"
	"github.com/swiggy-2022-bootcamp/training-i-plus-plus/nishant/mini-project/doctor_service/db"
	models "github.com/swiggy-2022-bootcamp/training-i-plus-plus/nishant/mini-project/doctor_service/model"
	"github.com/swiggy-2022-bootcamp/training-i-plus-plus/nishant/mini-project/doctor_service/producer"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

// CreateAppointment godoc
// @Summary create Appointment
// @Description create Appointment
// @Tags Appointment
// @Param   appointment     body entity.AppointmentRequest true  "appointment info"
// @Accept  json
// @Success 200
// @Failure 500
// @Router /appointment [post]
func CreateAppointment(c *gin.Context) {

	doctor := db.Session.DB(db.Mongo.Database).C(models.DoctorCollectionName)

	apreq := entity.AppointmentRequest{}

	if err := c.Bind(&apreq); err != nil {
		c.Error(err)
		return
	}

	toAdd := models.Appointment{
		Patient: apreq.Patient.UserId,
		From:    toTimeStamp(apreq.From),
		To:      toTimeStamp(apreq.To),
	}

	change := bson.M{"$push": bson.M{"appointments": toAdd}}

	log.Printf("Change : %+v", change)

	if err := doctor.UpdateId(bson.ObjectIdHex(apreq.Doctor), change); err != nil {
		c.Error(err)
		log.Printf("error while inserting appointment : %+v", err)
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	producer.Notifier.SendAppointmentEmail(apreq)
	c.Status(200)
}

type UserAppResponse struct {
	Doctor string
	From   string `example:"02 Jan 22 15:00 IST"`
	To     string `example:"02 Jan 22 16:00 IST"`
}

// GetAppointmentByUser godoc
// @Summary Get Appointment By User
// @Description Get Appointment By  User
// @Tags Appointment
// @Param   userid     path string true  "userid"
// @Accept  json
// @Success 200  {array} UserAppResponse
// @Failure 400
// @Failure 500
// @Router /appointment/user/{userid} [get]
func GetAppointmentByUser(c *gin.Context) {

	userid := c.Param("userid")
	if userid == "" {
		c.AbortWithStatusJSON(400, gin.H{
			"error": "userid not found",
		})
		return
	}

	doctor := db.Session.DB(db.Mongo.Database).C(models.DoctorCollectionName)
	result := []models.Doctor{}
	match := bson.M{"appointments.patient": userid}
	if err := doctor.Find(match).All(&result); err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	var res []UserAppResponse

	for _, doc := range result {
		for _, app := range doc.Appointments {
			res = append(res, UserAppResponse{
				Doctor: doc.Name,
				From:   tsToString(app.From),
				To:     tsToString(app.To),
			})
		}
	}

	c.JSON(200, &res)
}

func tsToString(i int64) string {
	return time.Unix(i, 0).Format(time.RFC822)
}

func toTimeStamp(d string) int64 {
	log.Println("Converting date time " + d)
	t, err := time.Parse(time.RFC822, d)
	if err != nil {
		log.Println("error while converting time " + err.Error())
	}
	log.Println(t.Unix())
	return t.Unix()
}
