package controllers

import (
	"context"
	//"errors"
	"fmt"
	"net/http"
	"sanitaria-microservices/doctorModule/configs"
	"sanitaria-microservices/doctorModule/models"
	"sanitaria-microservices/doctorModule/responses"
	"sanitaria-microservices/doctorModule/services"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var doctorCollection *mongo.Collection = configs.GetCollection(configs.DB, "doctors")
var validate = validator.New()

const (
    topic         = "Appointment"
)

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
		token, err := services.CreateToken(foundDoctor.User.EmailId, foundDoctor.User.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: token, Data: map[string]interface{}{"data": foundDoctor}})
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

		update := bson.M{"user.name": doctor.User.Name,"user.address": doctor.User.Address, "category": doctor.Category, "medicalLicenseLink": doctor.MedicalLicenseLink, "yoe": doctor.Yoe}

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
        id := c.Param("doctorId")
        var doctor models.Doctor
        defer cancel()

        objId, _ := primitive.ObjectIDFromHex(id)

        err := doctorCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&doctor)
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

		var appointment models.Appointment
		if err = c.BindJSON(&appointment); err!=nil{
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
		}
		newAppointment := models.Appointment{
			Id : primitive.NewObjectID(),
			Slot : appointment.Slot,
			Fees : appointment.Fees,
			Occupied : false,
			DoctorID : objId, 
		}
		currentAppointments := doctor.Appointments
		currentAppointments = append(currentAppointments,newAppointment)

		update := bson.M{"appointments": currentAppointments}

		result, updateErr := doctorCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

		if updateErr != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": updateErr.Error()}})
			return
		}
		p, err_ :=  configs.CreateProducer()
		if err_ != nil{
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err_.Error()}})
			return
		}
		configs.ProduceAppointment(p,topic,newAppointment)
		var updatedDoctor models.Doctor
		if result.MatchedCount == 1 {
			err = doctorCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedDoctor)

			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}
		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedDoctor}})
	}
}

// func BookAppointmentForGeneralUser(userId string) (models.Appointment, error){
// 	ctx,cancel := context.WithTimeout(context.Background(),10 * time.Second)
// 	defer cancel()

// 	results,err := doctorCollection.Find(ctx,bson.M{})
// 	if err != nil {
// 		return models.Appointment{}, errors.New(" No doctors available")
// 	}

// 	defer results.Close(ctx)
// 		for results.Next(ctx) {
// 			var singleDoctor models.Doctor
// 			if err = results.Decode(&singleDoctor); err != nil {
// 				return models.Appointment{}, errors.New("could not fetch doctor details")
// 			}

// 			allAppointments := singleDoctor.Appointments
			
// 			for i,appointment := range allAppointments{
// 				if !appointment.Occupied{
// 					appointment.Occupied = true;
// 					allAppointments[i] = appointment

// 					update := bson.M{"appointments":allAppointments}
// 					_, err := doctorCollection.UpdateOne(ctx, bson.M{"_id": singleDoctor.Id}, bson.M{"$set": update})
// 					if err != nil {
// 						return models.Appointment{}, errors.New("failed to update doctor details")
// 					}


// 					return appointment,nil
// 				}
// 			}

		
// 		}

// 		return models.Appointment{}, errors.New("no appointment available currently. please check back later")
// }

// func ListAvailableAppointments() ([]models.Appointment, error){
// 	ctx,cancel := context.WithTimeout(context.Background(),10 * time.Second)
// 	defer cancel()
// 	availableAppointments := make([]models.Appointment,0, 10)
// 	results,err := doctorCollection.Find(ctx,bson.M{})
// 	if err != nil {
// 		fmt.Println("no doctors available")
// 		return availableAppointments, errors.New("no doctors available")
// 	}
	
// 	defer results.Close(ctx)
// 		for results.Next(ctx) {
// 			var singleDoctor models.Doctor
// 			if err = results.Decode(&singleDoctor); err != nil {
// 				fmt.Println("could not fetch doctor details")
// 				return availableAppointments, errors.New("could not fetch doctor details")
// 			}

// 			allAppointments := singleDoctor.Appointments
			
// 			for _,appointment := range allAppointments{
// 				if !appointment.Occupied{
// 					availableAppointments = append(availableAppointments, appointment)
// 				}
// 			}

		
// 		}

// 		return availableAppointments, nil
// }