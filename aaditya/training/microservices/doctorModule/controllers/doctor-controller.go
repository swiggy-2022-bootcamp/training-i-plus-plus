package controllers

import (
	"context"
	"fmt"
	"net/http"
	"doctorModule/configs"
	"doctorModule/models"
	"doctorModule/responses"
	"doctorModule/services"
	"time"
	"log"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
    "github.com/segmentio/kafka-go"
)

const (
    topic         = "UserTopic"
    brokerAddress = "localhost:9092"
)

var doctorCollection *mongo.Collection = configs.GetCollection(configs.DB, "doctors")
var validate = validator.New()

var ctx = context.Background()
var l = log.New(os.Stdout, "kafka reader: ", 0)
var r = kafka.NewReader(kafka.ReaderConfig{
	Brokers: []string{brokerAddress},
	Topic:   topic,
	GroupID: "my-group",
	// assign the logger to the reader
	Logger: l,
})
func RegisterDoctor() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        var doctor models.Doctor
        var user models.User
        defer cancel()

        //validate the request body
        if err := c.BindJSON(&doctor); err != nil {
            c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }
        user = doctor.User

        //use the validator library to validate required fields
        if validationErr := validate.Struct(&doctor); validationErr != nil {
            c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
            return
        }
        if validationErr := validate.Struct(&user); validationErr != nil {
            c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
            return
        }
		hashPassword := services.HashPassword(user.Password)
		user.Password = hashPassword
        newDoctor := models.Doctor{
            Id:       primitive.NewObjectID(),
            Category:     doctor.Category,
            Yoe: 		  doctor.Yoe,
            MedicalLicenseLink:    doctor.MedicalLicenseLink,
			User:				   user,
        }
      
        result, err := doctorCollection.InsertOne(ctx, newDoctor)
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

        c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
    }
}

func LoginDoctor() gin.HandlerFunc {
	return func (c *gin.Context)  {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var doctor models.Doctor
		var foundDoctor models.Doctor

		if err := c.BindJSON(&doctor); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return 
		}
		fmt.Println(doctor.User.EmailId)
		err := doctorCollection.FindOne(ctx, bson.M{"user.emailid":doctor.User.EmailId}).Decode(&foundDoctor)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"email or password is incorrect"})
			return
		}

		passwordIsValid, msg := services.VerifyPassword(doctor.User.Password, foundDoctor.User.Password)
		defer cancel()
		if !passwordIsValid{
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		if foundDoctor.User.EmailId == ""{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"user not found"})
		}
		//token, err := services.CreateToken(foundDoctor.User.EmailId, foundDoctor.User.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "Success", Data: map[string]interface{}{"data": foundDoctor}})
	}
}

func GetDoctorByID() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        id := c.Param("id")
        var doctor models.Doctor
        defer cancel()

        objId, _ := primitive.ObjectIDFromHex(id)

        err := doctorCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&doctor)
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

        c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": doctor}})
    }
}

func EditDoctorByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		id := c.Param("id")
		var doctor models.Doctor
        var user models.User
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(id)

		//validate the request body
		if err := c.BindJSON(&doctor); err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

        user = doctor.User
		//use the validator library to validate required fields
		if validationErr := validate.Struct(&doctor); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

        if validationErr := validate.Struct(&user); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		update := bson.M{"user": doctor.User, "category": doctor.Category, "medicalLicenseLink": doctor.MedicalLicenseLink, "yoe": doctor.Yoe}

		result, err := doctorCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//get updated user details
		var updatedDoctor models.Doctor
		if result.MatchedCount == 1 {
			err := doctorCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedDoctor)

			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedDoctor}})
	}
}

func DeleteDoctorByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		id := c.Param("id")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(id)

		result, err := doctorCollection.DeleteOne(ctx, bson.M{"_id": objId})

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

func GetAllDoctors() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var doctors []models.Doctor
		defer cancel()

		results, err := doctorCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleDoctor models.Doctor
			if err = results.Decode(&singleDoctor); err != nil {
				c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			doctors = append(doctors, singleDoctor)
		}

		c.JSON(http.StatusOK,
			responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": doctors}},
		)
	}
}

func OpenSlotsForAppointments() gin.HandlerFunc{
	return func (c *gin.Context)  {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        id := c.Param("id")
        var doctor models.Doctor
        defer cancel()

        objId, _ := primitive.ObjectIDFromHex(id)

        err := doctorCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&doctor)
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

		var newAppointment models.Appointment
		if err = c.BindJSON(&newAppointment); err!=nil{
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
		}
		currentAppointments := doctor.Appointments
		currentAppointments = append(currentAppointments,newAppointment)

		update := bson.M{"appointments": currentAppointments}

		result, err := doctorCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//get updated user details
		var updatedDoctor models.Doctor
		if result.MatchedCount == 1 {
			err := doctorCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedDoctor)

			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedDoctor}})
	}
}

func GetAppointmentDetails() gin.HandlerFunc{
	
	return func (c *gin.Context)  {
		//ctx := context.Background()
		//appointment := consume(ctx)
		//appointments := make([]string,1)
		
		appointments := consume(ctx,r,l)
		c.JSON(http.StatusOK,
			responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": appointments}},
		)
	}
}

func consume(ctx context.Context, r *kafka.Reader, l *log.Logger) string{
    // create a new logger that outputs to stdout
    // and has the `kafka reader` prefix
    
    
        // the `ReadMessage` method blocks until we receive the next event
        msg, err := r.ReadMessage(ctx)
        if err != nil {
            panic("could not read message " + err.Error())
        }
        // after receiving the message, log its value
        fmt.Println("received: ", string(msg.Value))
		appointments := string(msg.Value)
		
		return appointments
}