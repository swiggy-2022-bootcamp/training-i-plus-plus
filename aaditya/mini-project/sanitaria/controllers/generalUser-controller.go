package controllers

import (
	"context"
	"net/http"
	"sanitaria/configs"
	"sanitaria/models"
	"sanitaria/responses"
	"sanitaria/services"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var generalUserCollection *mongo.Collection = configs.GetCollection(configs.DB, "generalUsers")
//var validate = validator.New()

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
            Id:       primitive.NewObjectID(),
            PreviousDiseases:     generalUser.PreviousDiseases,
            IsPatient: generalUser.IsPatient,
			User:				   user,
        }
      
        result, err := generalUserCollection.InsertOne(ctx, newGeneralUser)
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

        c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
    }
}

func LoginGeneralUser() gin.HandlerFunc {
	return func (c *gin.Context)  {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var generalUser models.GeneralUser
		var foundGenerealUser models.GeneralUser

		if err := c.BindJSON(&generalUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return 
		}
		err := generalUserCollection.FindOne(ctx, bson.M{"user.emailid":generalUser.User.EmailId}).Decode(&foundGenerealUser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"email or password is incorrect"})
			return
		}

		passwordIsValid, msg := services.VerifyPassword(generalUser.User.Password, foundGenerealUser.User.Password)
		defer cancel()
		if !passwordIsValid{
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		if foundGenerealUser.User.EmailId == ""{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"user not found"})
		}
		token, err := services.CreateToken(foundGenerealUser.User.EmailId, foundGenerealUser.User.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: token, Data: map[string]interface{}{"data": foundGenerealUser}})
	}
}

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

		update := bson.M{"user": generalUser.User, "ispatient": generalUser.IsPatient, "previousdisease": generalUser.PreviousDiseases}

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

func BookAppointment() gin.HandlerFunc{
	return func (c *gin.Context)  {
		ctx,cancel := context.WithTimeout(context.Background(),10 * time.Second)
		id := c.Param("id")
		defer cancel()

		var generalUser models.GeneralUser

		objId,_ := primitive.ObjectIDFromHex(id)

		err := generalUserCollection.FindOne(ctx,bson.M{"_id":objId}).Decode(&generalUser)
		if err != nil {
            c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

		scheduledAppointment,err_ := BookAppointmentForGeneralUser(id)
		if err_ != nil{
			c.JSON(http.StatusOK,responses.UserResponse{Status: http.StatusOK, Message: "Success", Data: map[string]interface{}{"data": err_}})
		}
		c.JSON(http.StatusOK,responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": scheduledAppointment}})
	}
}