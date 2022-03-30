package controllers

import (
	"context"
	"log"
	"net/http"
	"time"
	"userService/database"
	helper "userService/helpers"
	"userService/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(
	database.MongoClient,
	"user",
)
var validate *validator.Validate

func HashPassword() {

}

func VerifyPassword() {

}

// ShowAccount godoc
// @Summary      Sign up the user
// @Description  user registration API
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        FirstName  body 	string  true  "first name of user"
// @Param        LastName 	body	string  true  "last name of user"
// @Param        Password 	body	string  true  "password"
// @Param        Email 		body	string  true  "email id"
// @Success      200  {object} 	models.SignUpResponse
// @Failure      400  {number} 	http.StatusBadRequest
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /user/signup [post]
func Signup() gin.HandlerFunc {
	return func(g *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
		defer cancel()
		var user models.User

		if err := g.BindJSON(&user); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validate = validator.New()
		validationErr := validate.Struct(user)
		if validationErr != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		defer cancel()
		if err != nil {
			log.Panic(err)
			g.JSON(
				http.StatusBadRequest,
				gin.H{"error": "Error! occured while checking for email"},
			)
			return
		}

		if count > 0 {
			// log.Panic("Email is already exists.")
			g.JSON(http.StatusBadRequest, gin.H{"error": "User already exists."})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword(
			[]byte(user.Password),
			bcrypt.DefaultCost,
		)
		if err != nil {
			log.Panic("Password not hashed!")
			g.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}
		user.Password = string(hashedPassword)

		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Id = primitive.NewObjectID()
		user.User_id = user.Id.Hex()
		defer cancel()
		if err != nil {
			g.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		user.Token, err = helper.CreateToken(
			user.Email,
			user.FirstName,
			user.LastName,
			user.User_id,
		)
		defer cancel()
		if err != nil {
			g.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		insertedNumber, insertErr := userCollection.InsertOne(ctx, user)
		defer cancel()
		if insertErr != nil {
			g.JSON(http.StatusBadGateway, gin.H{"error": insertErr.Error()})
			return
		}

		g.JSON(
			http.StatusOK,
			gin.H{"insertedNumber": insertedNumber, "Token": user.Token},
		)
	}
}

// ShowAccount godoc
// @Summary      Login the user
// @Description  user log in API
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        Password 	body	string  true  "password"
// @Param        Email 		body	string  true  "email id"
// @Success      200  {object} 	models.User
// @Failure      400  {number} 	http.StatusBadRequest
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /user/login [post]
func Login() gin.HandlerFunc {
	return func(g *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()
		var user models.User

		if err := g.BindJSON(&user); err != nil {
			g.JSON(
				http.StatusBadRequest,
				gin.H{"error": "User JSON bind error", "details": err.Error()},
			)
			return
		}

		var userRecord models.User
		err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).
			Decode(&userRecord)
		defer cancel()
		if err != nil {
			g.JSON(
				http.StatusBadRequest,
				gin.H{"error": "User not found", "details": err.Error()},
			)
			return
		}

		err = bcrypt.CompareHashAndPassword(
			[]byte(userRecord.Password),
			[]byte(user.Password),
		)
		defer cancel()
		if err != nil {
			g.JSON(
				http.StatusBadGateway,
				gin.H{"error": "wrong password", "details": err.Error()},
			)
			return
		}

		g.JSON(http.StatusOK, gin.H{"user": userRecord})
	}
}

func GetUsers(gin *gin.Context) {

}

// ShowAccount 	 godoc
// @Summary      Get user details on ID
// @Description  get user details using ID
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        User_ID 	body	string  true  "unique user id"
// @Param        Token  	header	string  true  "user token"
// @Success      200  {object} 	models.User
// @Failure      400  {number} 	http.StatusBadRequest
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /user/getUserDetails [post]
func GetUserDetails() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
		defer cancel()

		var body struct {
			UserID string
		}

		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var UserDetails models.User
		err := userCollection.FindOne(ctx, bson.M{"user_id": body.UserID}).
			Decode(&UserDetails)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		response := struct {
			firstName string
			lastName  string
			Email     string
			Phone     string
			UserId    string
		}{
			UserDetails.FirstName,
			UserDetails.LastName,
			UserDetails.Email,
			UserDetails.Phone,
			UserDetails.User_id,
		}

		c.JSON(http.StatusOK, gin.H{"userDetails": response})
	}
}
