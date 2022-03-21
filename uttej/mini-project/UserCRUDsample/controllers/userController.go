package controllers

import (
	"log"

	"golang-app/models"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	UserController struct {
		session *mgo.Session
	}
)

const (
	DB_NAME       = "TrainTicketLelo"
	DB_COLLECTION = "users"
)

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func messageTypeDefault(msg string, c *gin.Context) {
	content := gin.H{
		"status": "200",
		"result": msg,
	}
	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(201, content)
}

func checkErrTypeOne(err error, msg string, status string, c *gin.Context) {
	if err != nil {
		panic(err)
		log.Fatalln(msg, err)
		content := gin.H{
			"status": status,
			"result": msg,
		}
		c.Writer.Header().Set("Content-Type", "application/json")
		c.JSON(200, content)

	}
}

func checkErrTypeTwo(msg string, status string, c *gin.Context) {
	content := gin.H{
		"status": status,
		"result": msg,
	}
	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(200, content)
}

func (uc UserController) UsersList(c *gin.Context) {

	var results []models.User
	err := uc.session.DB(DB_NAME).C(DB_COLLECTION).Find(nil).All(&results)
	if err != nil {
		checkErrTypeOne(err, "Users doesn't exist", "404", c)
		return
	}

	c.JSON(200, results)
}

func (uc UserController) GetUser(c *gin.Context) {
	id := c.Params.ByName("id")

	if !bson.IsObjectIdHex(id) {
		checkErrTypeTwo("ID is not a bson.ObjectId", "404", c)
		return
	}
	oid := bson.ObjectIdHex(id)

	u := models.User{}
	err := uc.session.DB(DB_NAME).C(DB_COLLECTION).FindId(oid).One(&u)

	if err != nil {
		checkErrTypeTwo("Users doesn't exist", "404", c)
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(200, u)
}

func (uc UserController) CreateUser(c *gin.Context) {

	var json models.User

	c.Bind(&json)

	u := uc.create_user(json.Name, json.Gender, json.Age, c)

	if u.Name == json.Name {
		content := gin.H{
			"result": "Success",
			"Name":   u.Name,
			"Gender": u.Gender,
			"Age":    u.Age,
		}

		c.Writer.Header().Set("Content-Type", "application/json")
		c.JSON(201, content)
	} else {
		c.JSON(500, gin.H{"result": "An error occured"})
	}

}

func (uc UserController) RemoveUser(c *gin.Context) {
	id := c.Params.ByName("id")

	if !bson.IsObjectIdHex(id) {
		checkErrTypeTwo("ID is not a bson.ObjectId", "404", c)
		return
	}
	oid := bson.ObjectIdHex(id)

	if err := uc.session.DB(DB_NAME).C(DB_COLLECTION).RemoveId(oid); err != nil {
		checkErrTypeOne(err, "Fail to Remove", "404", c)
		return
	}

	messageTypeDefault("Successfully Removed User", c)

}

func (uc UserController) UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var json models.User

	c.Bind(&json)

	if !bson.IsObjectIdHex(id) {
		checkErrTypeTwo("ID is not a bson.ObjectId", "404", c)
		return
	}

	u := uc.update_user(id, json.Name, json.Gender, json.Age, c)

	if u.Name == json.Name {
		content := gin.H{
			"result": "User Successfully Updated!",
			"Name":   u.Name,
			"Gender": u.Gender,
			"Age":    u.Age,
		}

		c.Writer.Header().Set("Content-Type", "application/json")
		c.JSON(201, content)
	} else {
		c.JSON(500, gin.H{"result": "An error occured"})
	}

}

func (uc UserController) create_user(Name string, Gender string, Age int, c *gin.Context) models.User {
	user := models.User{
		Name:   Name,
		Gender: Gender,
		Age:    Age,
	}
	err := uc.session.DB(DB_NAME).C(DB_COLLECTION).Insert(&user)
	checkErrTypeOne(err, "Insert failed", "403", c)
	return user
}

func (uc UserController) update_user(Id string, Name string, Gender string, Age int, c *gin.Context) models.User {

	user := models.User{

		Name:   Name,
		Gender: Gender,
		Age:    Age,
	}

	oid := bson.ObjectIdHex(Id)
	if err := uc.session.DB(DB_NAME).C(DB_COLLECTION).UpdateId(oid, &user); err != nil {
		checkErrTypeOne(err, "Update failed", "403", c)

	}

	return user
}
