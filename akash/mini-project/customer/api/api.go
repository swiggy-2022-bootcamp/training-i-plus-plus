package api

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"sample.akash.com/db"
	"sample.akash.com/log"
	"sample.akash.com/model"
)

var (
	repo db.CustomerRepository
)

func InitCustomerAPI(repository db.CustomerRepository) {
	repo = repository
	repository.Connect()
}

func Login(c *gin.Context) {

	loginData := model.LoginData{}
	if err := c.BindJSON(&loginData); err != nil {
		panic(err)
	}
	log.Info(loginData)

	user := repo.FindOneWithUsername(loginData.Username)
	if user != nil && arePasswordsEqual(user.Password, loginData.Password) == true {
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

	user := repo.FindOneWithUsername(userData.Username)
	if user == nil {
		userData.Password = hashPassword(userData.Password)
		repo.SaveUser(userData)
		c.Data(http.StatusOK, "application/json", []byte(`{"message":"register successful"}`))
	} else {
		c.Data(http.StatusUnauthorized, "application/json", []byte(`{"message":"user already exist with this username"}`))
	}
}

func QueryOne(c *gin.Context) {

	username := c.Param("username")
	log.Info("Find user with username : ", username)

	user := repo.FindOneWithUsername(username)
	if user != nil {
		c.JSON(http.StatusOK, user)
	} else {
		c.Data(http.StatusUnauthorized, "application/json", []byte(`{"message":"user not found"}`))
	}
}

func QueryAll(c *gin.Context) {
	users := repo.FindAll()
	c.JSON(http.StatusOK, users)
}

func Delete(c *gin.Context) {

	username := c.Param("username")
	log.Info("Delete user with username : ", username)

	if repo.DeleteUser(username) == true {
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

	if repo.FindAndUpdate(userData) == true {
		c.Data(http.StatusOK, "application/json", []byte(`{"message":"update successful"}`))
	} else {
		c.Data(http.StatusUnauthorized, "application/json", []byte(`{"message":"update failed"}`))
	}
}

func hashPassword(userPassword string) string {

	password := []byte(userPassword)

	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		log.Error(err)
	}
	log.Info(string(hashedPassword))

	return string(hashedPassword[:])
}

func arePasswordsEqual(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
