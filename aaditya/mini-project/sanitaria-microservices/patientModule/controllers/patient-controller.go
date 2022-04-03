package controllers

import (
	"context"
	"net/http"
	"sanitaria-microservices/patientModule/configs"
	"sanitaria-microservices/patientModule/models"
	"sanitaria-microservices/patientModule/responses"
	"sanitaria-microservices/patientModule/services"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var patientCollection *mongo.Collection = configs.GetCollection(configs.DB, "patients")
var validate = validator.New()

// RegisterPatient godoc
// @Summary To register a new patient in the sanitaria application
// @Description This request will create a new patient profile for a user.
// @Tags Patient
// @Schemes
// @Accept json
// @Produce json
// @Param req body models.Patient true "Patient details"
// @Success	201  {object} 	models.Patient
// @Failure	400  {number} 	http.http.StatusBadRequest
// @Failure	500  {number} 	http.StatusInternalServerError
// @Router /patientRegistration [POST]
func RegisterPatient() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var patient models.Patient
		var user models.User
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&patient); err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		user = patient.User

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&patient); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}
		if validationErr := validate.Struct(&user); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		hashPassword := services.HashPassword(user.Password)
		user.Password = hashPassword

		newPatient := models.Patient{
			Id:               primitive.NewObjectID(),
			DoctorAssignedId: patient.DoctorAssignedId,
			IsDischarged:     patient.IsDischarged,
			RoomAllocated:    patient.RoomAllocated,
			User:             user,
		}

		result, err := patientCollection.InsertOne(ctx, newPatient)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

// LoginPatient godoc
// @Summary User login for a patient profile.
// @Description This request will login a patient.
// @Tags Patient
// @Schemes
// @Accept json
// @Produce json
// @Param req body models.User true "User emailid and password"
// @Success	200  {string} 	token
// @Failure	400  {number} 	http.http.StatusBadRequest
// @Failure	404  {number} 	http.http.StatusNotFound
// @Failure	500  {number} 	http.StatusInternalServerError
// @Router /patientLogin [POST]
func LoginPatient() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var patient models.Patient
		var foundPatient models.Patient

		if err := c.BindJSON(&patient); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := patientCollection.FindOne(ctx, bson.M{"user.emailid": patient.User.EmailId}).Decode(&foundPatient)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "email or password is incorrect"})
			return
		}

		passwordIsValid, msg := services.VerifyPassword(patient.User.Password, foundPatient.User.Password)
		defer cancel()
		if !passwordIsValid {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		if foundPatient.User.EmailId == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
		}
		token, err := services.CreateToken(foundPatient.User.EmailId, foundPatient.User.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: token, Data: map[string]interface{}{"data": foundPatient}})
	}
}

// GetPatientByID godoc
// @Summary Get patient by ID.
// @Description View all the details of a patient.
// @Tags Patient
// @Schemes
// @Param id path string true "Patient id"
// @Accept json
// @Produce json
// @Success	200  {object} 	models.Patient
// @Failure	500  {number} 	http.StatusInternalServerError
// @Security Bearer Token
// @Router /patient/{id} [GET]
func GetPatientByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		id := c.Param("id")
		var patient models.Patient
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(id)

		err := patientCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&patient)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": patient}})
	}
}

// EditPatientByID godoc
// @Summary Edit patient by ID.
// @Description Edit details of a patient.
// @Tags Patient
// @Schemes
// @Param id path string true "Patient id"
// @Accept json
// @Produce json
// @Param req body models.Patient true "Patient details"
// @Success	200  {object} 	models.Patient
// @Failure	400  {number} 	http.StatusBadRequest
// @Failure	500  {number} 	http.StatusInternalServerError
// @Security Bearer Token
// @Router /patient/{id} [PUT]
func EditPatientByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		id := c.Param("id")
		var patient models.Patient
		var user models.User
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(id)

		//validate the request body
		if err := c.BindJSON(&patient); err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		user = patient.User
		//use the validator library to validate required fields
		if validationErr := validate.Struct(&patient); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		if validationErr := validate.Struct(&user); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		update := bson.M{"user.name": patient.User.Name, "user.address": patient.User.Address, "doctorassignedid": patient.DoctorAssignedId, "isdischarged": patient.IsDischarged, "roomallocated": patient.RoomAllocated}

		result, err := patientCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//get updated user details
		var updatedPatient models.Patient
		if result.MatchedCount == 1 {
			err := patientCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedPatient)

			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedPatient}})
	}
}

// DeletePatientByID godoc
// @Summary Delete patient by ID.
// @Description Delete a patient.
// @Tags Patient
// @Schemes
// @Param id path string true "Patient id"
// @Accept json
// @Produce json
// @Success	200  {string} 	User successfully deleted!
// @Failure	404  {number} 	http.http.StatusNotFound
// @Failure	500  {number} 	http.StatusInternalServerError
// @Security Bearer Token
// @Router /patient/{id} [DELETE]
func DeletePatientByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		id := c.Param("id")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(id)

		result, err := patientCollection.DeleteOne(ctx, bson.M{"_id": objId})

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

// GetAllPatients godoc
// @Summary Get all patients list.
// @Description Get details of all patients.
// @Tags Patient
// @Schemes
// @Accept json
// @Produce json
// @Success	200  {array} 	models.Patient
// @Failure	500  {number} 	http.StatusInternalServerError
// @Security Bearer Token
// @Router /patients [GET]
func GetAllPatients() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var patients []models.Patient
		defer cancel()

		results, err := patientCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singlePatient models.Patient
			if err = results.Decode(&singlePatient); err != nil {
				c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			patients = append(patients, singlePatient)
		}

		c.JSON(http.StatusOK,
			responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": patients}},
		)
	}
}
