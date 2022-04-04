package controllers

import (
	"context"
	"net/http"
	"time"
	"user_service/configs"
	"user_service/middlewares"
	"user_service/models"
	"user_service/responses"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

func GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			var users []models.User
			defer cancel()
		
			results, err := userCollection.Find(ctx, bson.M{})
		
			if err != nil {
					c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
					return
			}
		
			defer results.Close(ctx)
			for results.Next(ctx) {
					var singleUser models.User
					if err = results.Decode(&singleUser); err != nil {
							c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
					}
				
					users = append(users, singleUser)
			}
		
			c.JSON(http.StatusOK,
					responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": users}},
			)
	}
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			var user models.User
			defer cancel()
		
			//validate the request body
			if err := c.BindJSON(&user); err != nil {
					c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
					return
			}
		
			//use the validator library to validate required fields
			if validationErr := validate.Struct(&user); validationErr != nil {
					c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
					return
			}
		
		hashedPassword, err := middlewares.HashPassword(user.Password)

		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

			newUser := models.User{
					Id:       primitive.NewObjectID(),
					Name:     user.Name,
					Password: hashedPassword,
					Contact:  models.ContactInfo{
						Email: user.Contact.Email,
						Phone: user.Contact.Phone,
					},
					DateOfBirth:    user.DateOfBirth,
					IdProof:    user.IdProof,
					Role:    models.Customer,
			}
				
			result, err := userCollection.InsertOne(ctx, newUser)
			if err != nil {
					c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
					return
			}
		
			c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetUserById() gin.HandlerFunc {
	return func(c *gin.Context) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			userId := c.Param("userId")
			var user models.User
			defer cancel()
		
			objId, _ := primitive.ObjectIDFromHex(userId)
		
			err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)
			if err != nil {
					c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
					return
			}
			c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": user}})
	}
}

func UpdateUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        userId := c.Param("userId")
        var user models.User
        defer cancel()
      
        objId, _ := primitive.ObjectIDFromHex(userId)
      
        //validate the request body
        if err := c.BindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }
      
        //use the validator library to validate required fields
        if validationErr := validate.Struct(&user); validationErr != nil {
            c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
            return
        }
      
        update := bson.M{
					"name": user.Name, 
					"password": user.Password, 
					"contact": user.Contact,
					"DateOfBirth": user.DateOfBirth,
					"idProof": user.IdProof,
				}
        result, err := userCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
      
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }
      
        //get updated user details
        var updatedUser models.User
        if result.MatchedCount == 1 {
            err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedUser)
            if err != nil {
                c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
                return
            }
        }
      
        c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedUser}})
    }
}

func DeleteUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        userId := c.Param("userId")
        defer cancel()
      
        objId, _ := primitive.ObjectIDFromHex(userId)
      
        result, err := userCollection.DeleteOne(ctx, bson.M{"id": objId})
      
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


func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User
		var regUser models.User
		defer cancel()

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error()})
			return
		}

		if err := userCollection.FindOne(ctx, bson.M{"name": user.Name}).Decode(&regUser); err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "Not Registered", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

	  err := middlewares.CheckPassword(regUser.Password, user.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "Wrong Password", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		
		if regUser.Role == 0 {

			token, err := middlewares.GetJWT("ADMIN", regUser.Name, regUser.Contact.Email, regUser.Id)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error in generating token", Data: map[string]interface{}{"data": err.Error()}})
				return
			}

			c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "Logged In successfully",Data: map[string]interface{}{"token" : token}})
			return

		} else if regUser.Role == 1 {

			token, err := middlewares.GetJWT("USER", regUser.Name, regUser.Contact.Email, regUser.Id)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error in generating token", Data: map[string]interface{}{"data": err.Error()}})
				return
			}

			c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "Logged In successfully",Data: map[string]interface{}{"token" : token}})
			return

		} else {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusInternalServerError, Message: "Role cannot be determined"})
			return
		}
	}
}