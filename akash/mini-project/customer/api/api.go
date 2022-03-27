package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sample.akash.com/db"
	"sample.akash.com/log"
	"sample.akash.com/model"
)

func Login(c *gin.Context) {

	loginData := model.LoginData{}
	if err := c.BindJSON(&loginData); err != nil {
		panic(err)
	}
	log.Info(loginData)

	user := db.FindOneWithUsername(loginData.Username)
	if user != nil && loginData.Password == user.Password {
		c.Data(http.StatusOK, "application/json", []byte(`{"message":"login successful"}`))
	} else {
		c.Data(http.StatusUnauthorized, "application/json", []byte(`{"message":"invalid credentials"}`))
	}
}

func Register(c *gin.Context) {

	userData := model.User{}
	if err := c.BindJSON(&userData); err != nil {
		panic(err)
	}

	log.Info("Register request for user : ", userData)

	user := db.FindOneWithUsername(userData.Username)
	if user == nil {
		db.SaveUser(userData)
		c.Data(http.StatusOK, "application/json", []byte(`{"message":"register successful"}`))
	} else {
		c.Data(http.StatusUnauthorized, "application/json", []byte(`{"message":"user already exist with this username"}`))
	}
}

func QueryOne(c *gin.Context) {

	jsonData := struct {
		Username string `json:"username"`
	}{}
	if err := c.BindJSON(&jsonData); err != nil {
		panic(err)
	}
	log.Info("Find user with username : ", jsonData.Username)

	user := db.FindOneWithUsername(jsonData.Username)
	if user != nil {
		c.JSON(http.StatusOK, user)
	} else {
		c.Data(http.StatusUnauthorized, "application/json", []byte(`{"message":"user not found"}`))
	}
}

func QueryAll(c *gin.Context) {
	users := db.FindAll()
	c.JSON(http.StatusOK, users)
}

func Delete(c *gin.Context) {

	jsonData := struct {
		Username string `json:"username"`
	}{}
	if err := c.BindJSON(&jsonData); err != nil {
		panic(err)
	}
	log.Info("Delete user with username : ", jsonData.Username)

	if db.DeleteUser(jsonData.Username) == true {
		c.Data(http.StatusOK, "application/json", []byte(`{"message":"user delete successful"}`))
	} else {
		c.Data(http.StatusUnauthorized, "application/json", []byte(`{"message":"delete failed"}`))
	}

}

func Update(c *gin.Context) {

	userData := model.User{}
	if err := c.BindJSON(&userData); err != nil {
		panic(err)
	}

	log.Info("Update request for user : ", userData)

	if db.FindAndUpdate(userData) == true {
		c.Data(http.StatusOK, "application/json", []byte(`{"message":"update successful"}`))
	} else {
		c.Data(http.StatusUnauthorized, "application/json", []byte(`{"message":"update failed"}`))
	}
}
