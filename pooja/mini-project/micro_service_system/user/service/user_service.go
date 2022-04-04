package service

import (
	"context"
	"net/http"
	"time"
	"user/database"
	"user/helper"
	"user/model"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.GetCollection(database.DB, "users")
var validate *validator.Validate = validator.New()

// ShowAccount godoc
// @Summary      Regsiter user
// @Description  User regisration API
// @Tags         signup
// @Accept       json
// @Produce      json
// @Param        ID  			body 	primitive.ObjectID  true  "unique id for every user, auto-generated"
// @Param        UserName 		body	string   	true  "unique username for each user, provided by the user itself"
// @Param        EmailId 		body	string  true  "user's email address"
// @Param        Password  			body 	string  true  "user's account password for the system"
// @Success      201  {number}  http.StatusCreated
// @Failure      400  {number} 	http.StatusBadRequest
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /signup [post]
func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
		defer cancel()
		var user model.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error()})

			return
		}
		log.Info("new user signup", user)

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&user); validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": validationErr.Error()})
			log.Error("User information is not valid")
			return
		}
		var userFromDB model.User
		count, err := userCollection.CountDocuments(ctx, bson.M{"username": userFromDB.UserName})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error! occured while checking for email"})
			log.Error("Error occured while checking for email")
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists."})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		newUser := model.User{
			ID:       primitive.NewObjectID(),
			UserName: user.UserName,
			EmailId:  user.EmailId,
			Password: string(hashedPassword),
		}
		result, err := userCollection.InsertOne(ctx, newUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error()})
			log.Error("Error in adding user information to the db")
			return
		}
		c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "success", "data": result})
	}
}

// ShowAccount godoc
// @Summary      Login user
// @Description  User login API
// @Tags         login
// @Accept       json
// @Produce      json
// @Param        UserName 		body	string   	true  "unique username for each user, provided by the user itself"
// @Param        Password  			body 	string  true  "user's account password for the system"
// @Success      200  {string}  jwt.token
// @Failure      400  {number} 	http.StatusBadRequest
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /login [post]
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var login model.LoginDTO
		var user model.User
		defer cancel()
		log.Info("in login")
		if err := c.BindJSON(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error()})
			return
		}
		if err := userCollection.FindOne(ctx, bson.M{"username": login.Username}).Decode(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "user not registered", "error": err.Error()})
			return
		}
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": "incorrect password", "details": err.Error()})
			log.Info("User information is not valid")
			return
		}
		token, err := helper.CreateToken(user.UserName, user.EmailId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}
