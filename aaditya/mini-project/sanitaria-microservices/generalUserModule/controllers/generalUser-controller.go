package controllers

import (
	"context"
	"net/http"
	"sanitaria-microservices/generalUserModule/configs"
	"sanitaria-microservices/generalUserModule/models"
	"sanitaria-microservices/generalUserModule/responses"
	"sanitaria-microservices/generalUserModule/services"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var generalUserCollection *mongo.Collection = configs.GetCollection(configs.DB, "generalUsers")
var validate = validator.New()

// RegisterGeneralUser godoc
// @Summary To register a new generalUser in the sanitaria application
// @Description This request will create a new generalUser profile for a user.
// @Tags GeneralUser
// @Schemes
// @Accept json
// @Produce json
// @Param req body models.GeneralUser true "General user details"
// @Success	201  {object} 	models.GeneralUser
// @Failure	400  {number} 	http.http.StatusBadRequest
// @Failure	500  {number} 	http.StatusInternalServerError
// @Router /generalUserRegistration [POST]
func RegisterGeneralUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var generalUser models.GeneralUser
		var user models.User
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&generalUser); err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		user = generalUser.User

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&generalUser); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}
		if validationErr := validate.Struct(&user); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		hashPassword := services.HashPassword(user.Password)
		user.Password = hashPassword

		newGeneralUser := models.GeneralUser{
			Id:               primitive.NewObjectID(),
			PreviousDiseases: generalUser.PreviousDiseases,
			IsPatient:        generalUser.IsPatient,
			User:             user,
		}

		result, err := generalUserCollection.InsertOne(ctx, newGeneralUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

// LoginGeneralUser godoc
// @Summary User login for a generalUser profile.
// @Description This request will login a generalUser.
// @Tags GeneralUser
// @Schemes
// @Accept json
// @Produce json
// @Param req body models.User true "User emailid and password"
// @Success	200  {string} 	token
// @Failure	400  {number} 	http.http.StatusBadRequest
// @Failure	404  {number} 	http.http.StatusNotFound
// @Failure	500  {number} 	http.StatusInternalServerError
// @Router /generalUserLogin [POST]
func LoginGeneralUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var generalUser models.GeneralUser
		var foundGenerealUser models.GeneralUser

		if err := c.BindJSON(&generalUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := generalUserCollection.FindOne(ctx, bson.M{"user.emailid": generalUser.User.EmailId}).Decode(&foundGenerealUser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "email or password is incorrect"})
			return
		}

		passwordIsValid, msg := services.VerifyPassword(generalUser.User.Password, foundGenerealUser.User.Password)
		defer cancel()
		if !passwordIsValid {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		if foundGenerealUser.User.EmailId == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
		}
		token, err := services.CreateToken(foundGenerealUser.User.EmailId, foundGenerealUser.User.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: token, Data: map[string]interface{}{"data": foundGenerealUser}})
	}
}

// GetGeneralUserByID godoc
// @Summary Get generalUser by ID.
// @Description View all the details of a generalUser.
// @Tags GeneralUser
// @Schemes
// @Param id path string true "GeneralUser id"
// @Accept json
// @Produce json
// @Success	200  {object} 	models.GeneralUser
// @Failure	500  {number} 	http.StatusInternalServerError
// @Security Bearer Token
// @Router /generalUser/{id} [GET]
func GetGeneralUserByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		id := c.Param("id")
		var generalUser models.GeneralUser
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(id)

		err := generalUserCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&generalUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": generalUser}})
	}
}

// EditGeneralUserByID godoc
// @Summary Edit generalUser by ID.
// @Description Edit details of a generalUser.
// @Tags GeneralUser
// @Schemes
// @Param id path string true "GeneralUser id"
// @Accept json
// @Produce json
// @Param req body models.GeneralUser true "General user details"
// @Success	200  {object} 	models.GeneralUser
// @Failure	400  {number} 	http.StatusBadRequest
// @Failure	500  {number} 	http.StatusInternalServerError
// @Security Bearer Token
// @Router /generalUser/{id} [PUT]
func EditGeneralUserByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		id := c.Param("id")
		var generalUser models.GeneralUser
		var user models.User
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(id)

		//validate the request body
		if err := c.BindJSON(&generalUser); err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		user = generalUser.User
		//use the validator library to validate required fields
		if validationErr := validate.Struct(&generalUser); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		if validationErr := validate.Struct(&user); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		update := bson.M{"user.name": generalUser.User.Name, "user.address": generalUser.User.Address, "ispatient": generalUser.IsPatient, "previousdiseases": generalUser.PreviousDiseases}

		result, err := generalUserCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//get updated user details
		var updatedGeneralUser models.GeneralUser
		if result.MatchedCount == 1 {
			err := generalUserCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedGeneralUser)

			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedGeneralUser}})
	}
}

// DeleteGeneralUserByID godoc
// @Summary Delete generalUser by ID.
// @Description Delete a generalUser.
// @Tags GeneralUser
// @Schemes
// @Param id path string true "GeneralUser id"
// @Accept json
// @Produce json
// @Success	200  {string} 	User successfully deleted!
// @Failure	404  {number} 	http.http.StatusNotFound
// @Failure	500  {number} 	http.StatusInternalServerError
// @Security Bearer Token
// @Router /generalUser/{id} [DELETE]
func DeleteGeneralUserByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		id := c.Param("id")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(id)

		result, err := generalUserCollection.DeleteOne(ctx, bson.M{"_id": objId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.UserResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "User with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "User successfully deleted!"}},
		)
	}
}

// GetAllGeneralUsers godoc
// @Summary Get all generalUsers list.
// @Description Get details of all generalUsers.
// @Tags GeneralUser
// @Schemes
// @Accept json
// @Produce json
// @Success	200  {array} 	models.GeneralUser
// @Failure	500  {number} 	http.StatusInternalServerError
// @Security Bearer Token
// @Router /generalUsers [GET]
func GetAllGeneralUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var generalUsers []models.GeneralUser
		defer cancel()

		results, err := generalUserCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleGeneralUser models.GeneralUser
			if err = results.Decode(&singleGeneralUser); err != nil {
				c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			generalUsers = append(generalUsers, singleGeneralUser)
		}

		c.JSON(http.StatusOK,
			responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": generalUsers}},
		)
	}
}
