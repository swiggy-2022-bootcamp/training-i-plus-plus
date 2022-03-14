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

	//TODO: Check if exist in DB
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
	log.Info(userData)

	//TODO: Check if exist in DB

	//TODO: Save user in DB
	SaveUser(userData)

	c.Data(http.StatusOK, "application/json", []byte(`{"message":"register successful"}`))
}

func QueryAll(c *gin.Context) {
	//TODO: Check if exist in DB
	users := FindAll()

	c.JSON(http.StatusOK, users)
}
