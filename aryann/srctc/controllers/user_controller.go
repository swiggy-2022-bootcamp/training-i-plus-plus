package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"srctc/config"
	"srctc/models"
	"srctc/responses"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = config.GetCollection(config.DB, "users")
var purchasedCollection *mongo.Collection = config.GetCollection(config.DB, "purchased")
var validate = validator.New()

const (
	topic         = "purchased"
	brokerAddress = "localhost:9092"
)

type kafka_booking_ticket struct {
	insertedid string
	purchased  models.Purchased
}

func init() {
	// go consume_booked_ticket()
}

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User
		defer cancel()

		// validate the json body
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// validate the user
		if validationErr := validate.Struct(&user); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newUser := models.User{
			Name:        user.Name,
			Email:       user.Email,
			PurchasedID: []primitive.ObjectID{},
		}

		result, err := userCollection.InsertOne(ctx, newUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		var user models.User
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		err := userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": user}})
	}
}

func EditUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userid")
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

		update := bson.M{"name": user.Name, "email": user.Email}
		result, err := userCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//get updated user details
		var updatedUser models.User
		if result.MatchedCount == 1 {
			err := userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedUser)
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
		userId := c.Param("userid")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		result, err := userCollection.DeleteOne(ctx, bson.M{"_id": objId})

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

func PurchaseTicket() gin.HandlerFunc {
	return func(c *gin.Context) {
		// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		// var purchased models.Purchased
		// defer cancel()

		// //validate the request body
		// if err := c.BindJSON(&purchased); err != nil {
		// 	c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error in binding", Data: map[string]interface{}{"data": err.Error()}})
		// 	return
		// }

		// //use the validator library to validate required fields
		// if validationErr := validate.Struct(&purchased); validationErr != nil {
		// 	c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error in validating", Data: map[string]interface{}{"data": validationErr.Error()}})
		// 	return
		// }

		//check and update avaiable tickets
		// var ticket models.Ticket

		// err := ticketCollection.FindOne(ctx, bson.M{"train_id": purchased.Train_id}).Decode(&ticket)
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, responses.PurchasedResponse{Status: http.StatusInternalServerError, Message: "Incorrect train id", Data: map[string]interface{}{"data": err.Error()}})
		// 	return
		// }

		// if ticket.Capacity == 0 {
		// 	c.JSON(http.StatusInternalServerError, responses.PurchasedResponse{Status: http.StatusInternalServerError, Message: "No tickets available", Data: map[string]interface{}{"data": err.Error()}})
		// 	return
		// }

		// update := bson.M{"capacity": ticket.Capacity - 1}
		// _, err = ticketCollection.UpdateOne(ctx, bson.M{"trainid": purchased.Train_id}, bson.M{"$set": update})

		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, responses.TicketResponse{Status: http.StatusInternalServerError, Message: "error in updating capacity", Data: map[string]interface{}{"data": err.Error()}})
		// 	return
		// }

		// var trainbooked models.Train

		// err = trainCollection.FindOne(ctx, bson.M{"_id": purchased.Train_id}).Decode(&trainbooked)

		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, responses.PurchasedResponse{Status: http.StatusInternalServerError, Message: "error in train find", Data: map[string]interface{}{"data": err.Error()}})
		// 	return
		// }

		// newpurchased := models.Purchased{
		// 	Train_id:        purchased.Train_id,
		// 	User_id:         purchased.User_id,
		// 	Departure:       trainbooked.Source,
		// 	Arrival:         trainbooked.Destination,
		// 	Departure_time:  ticket.Departure_time,
		// 	Arrival_time:    ticket.Arrival_time,
		// 	Passengers_info: purchased.Passengers_info,
		// }

		// result, err := purchasedCollection.InsertOne(ctx, newpurchased)
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		// 	return
		// }

		// iid := fmt.Sprintf("%v", result.InsertedID)
		// new_produce_ticket := kafka_booking_ticket{
		// 	insertedid: iid,
		// 	purchased:  newpurchased,
		// }

		// go produce_booked_ticket(new_produce_ticket)

		// c.JSON(http.StatusCreated, responses.PurchasedResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

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

		//reading from the db in an optimal way
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

func GetPurchased() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		purchasedId := c.Param("purchasedid")
		var purchased models.Purchased
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(purchasedId)

		err := purchasedCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&purchased)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PurchasedResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.PurchasedResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": purchased}})
	}
}

func DeletePurchased() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		purchasedId := c.Param("purchasedid")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(purchasedId)

		result, err := purchasedCollection.DeleteOne(ctx, bson.M{"_id": objId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PurchasedResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.PurchasedResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Purchased with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.PurchasedResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Purchased successfully deleted!"}},
		)
	}
}

func produce_booked_ticket(nbt kafka_booking_ticket) {
	l := log.New(os.Stdout, "kafka producer lol ", 0)
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		Logger:  l,
	})

	bytes, _ := json.Marshal(nbt.purchased)
	err := w.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(nbt.insertedid),
		Value: []byte(bytes),
	})
	if err != nil {
		panic("could not write message " + err.Error())
	}
}

func consume_booked_ticket() {
	l := log.New(os.Stdout, "kafka producer lol 2 ", 0)
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		Logger:  l,
	})

	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			panic("could not read message " + err.Error())
		}

		// after receiving the message, log its value
		fmt.Println("received: ", string(msg.Value))
		nbtr := models.Purchased{}
		json.Unmarshal([]byte(msg.Value), &nbtr)
		fmt.Println(nbtr)
		update_tickets_booked := bson.M{"tickets_booked": string(msg.Key)}
		_, err = userCollection.UpdateOne(context.Background(), bson.M{"_id": nbtr.User_id}, bson.M{"$push": update_tickets_booked})
		if err != nil {
			panic("could not update booked ticket " + err.Error())
		}
	}
}
