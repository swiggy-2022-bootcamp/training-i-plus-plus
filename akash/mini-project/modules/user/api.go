package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sample.akash.com/log"
	"sample.akash.com/model"
)

func Login(c *gin.Context) {

	loginData := model.LoginData{}
	if err := c.BindJSON(&loginData); err != nil {
		panic(err)
	}
	log.Info(loginData)

	user := FindOneWithEmail(loginData.Email)
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

	user := FindOneWithEmail(userData.Email)
	if user == nil {
		SaveUser(userData)
		c.Data(http.StatusOK, "application/json", []byte(`{"message":"register successful"}`))
	} else {
		c.Data(http.StatusUnauthorized, "application/json", []byte(`{"message":"user already exist with this email"}`))
	}
}

func QueryOne(c *gin.Context) {

	jsonData := struct {
		Email string `json:"email"`
	}{}
	if err := c.BindJSON(&jsonData); err != nil {
		panic(err)
	}
	log.Info("Find user with email : ", jsonData.Email)

	user := FindOneWithEmail(jsonData.Email)
	if user != nil {
		c.JSON(http.StatusOK, user)
	} else {
		c.Data(http.StatusUnauthorized, "application/json", []byte(`{"message":"user not found"}`))
	}
}

func QueryAll(c *gin.Context) {
	users := FindAll()
	c.JSON(http.StatusOK, users)
}

func Delete(c *gin.Context) {

	jsonData := struct {
		Email string `json:"email"`
	}{}
	if err := c.BindJSON(&jsonData); err != nil {
		panic(err)
	}
	log.Info("Delete user with email : ", jsonData.Email)

	if DeleteUser(jsonData.Email) == true {
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

	if FindAndUpdate(userData) == true {
		c.Data(http.StatusOK, "application/json", []byte(`{"message":"update successful"}`))
	} else {
		c.Data(http.StatusUnauthorized, "application/json", []byte(`{"message":"update failed"}`))
	}
}
