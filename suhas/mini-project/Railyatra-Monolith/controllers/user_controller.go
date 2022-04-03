package controllers

import (
	"context"
	"gin-mongo-api/kafka"
	"gin-mongo-api/models"
	"gin-mongo-api/repository"
	"gin-mongo-api/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//var userCollection *mongo.Collection = config.GetCollection(config.DB, "users")
//var bookedticketCollection *mongo.Collection = config.GetCollection(config.DB, "bookedtickets")
var validate = validator.New()

func init() {
	go kafka.Consume_avail_ticket()
	go kafka.Consume_train()
}

var userrepo repository.UserRepository
var bookedticketrepo repository.BookedTicketRepository

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
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

		newUser := models.User{
			Name:           user.Name,
			Email:          user.Email,
			BookedTicketID: []primitive.ObjectID{},
		}

		//result, err := userCollection.InsertOne(ctx, newUser)
		result, err := userrepo.Insert(newUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
		//new way
		//result,err := userrepo.Insert(newUser)
	}
}

func GetAUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userid")
		var user models.User
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		//err := userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)
		user, err := userrepo.Read(objId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": user}})
	}
}

func EditAUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
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

		updatedid, err := userrepo.Update(user, objId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"Updated id": updatedid}})
	}
}

func DeleteAUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userid")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		resultcount, err := userrepo.Delete(objId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if resultcount.(int) < 1 {
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

func GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var users []models.User
		defer cancel()

		users, err := userrepo.ReadAll()

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK,
			responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": users}},
		)
	}
}

func CreateBookedTicket() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var bookedticket models.BookedTicket
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&bookedticket); err != nil {
			c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error in binding", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := avalidate.Struct(&bookedticket); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error in validating", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		//check and update avaiable tickets
		var availticket models.AvailTicket

		//err := availticketCollection.FindOne(ctx, bson.M{"train_id": bookedticket.Train_id}).Decode(&availticket)
		availticket, err := availticketrepo.ReadTrainId(bookedticket.Train_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BookedTicketResponse{Status: http.StatusInternalServerError, Message: "Incorrect train id", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if availticket.Capacity == 0 {
			c.JSON(http.StatusInternalServerError, responses.BookedTicketResponse{Status: http.StatusInternalServerError, Message: "No tickets available", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// use this using kafka in future

		//update := bson.M{"capacity": availticket.Capacity - 1}
		//, err = availticketCollection.UpdateOne(ctx, bson.M{"trainid": bookedticket.Train_id}, bson.M{"$set": update})
		availticket.Capacity -= 1
		_, err = availticketrepo.Update(availticket, availticket.Id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AvailTicketResponse{Status: http.StatusInternalServerError, Message: "error in updating capacity", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		var trainbooked models.Train

		//err = trainCollection.FindOne(ctx, bson.M{"_id": bookedticket.Train_id}).Decode(&trainbooked)
		trainbooked, err = trainrepo.Read(bookedticket.Train_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BookedTicketResponse{Status: http.StatusInternalServerError, Message: "error in train find", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		newBookedTicket := models.BookedTicket{
			Train_id:        bookedticket.Train_id,
			User_id:         bookedticket.User_id,
			Departure:       trainbooked.Station1,
			Arrival:         trainbooked.Station2,
			Departure_time:  availticket.Departure_time,
			Arrival_time:    availticket.Arrival_time,
			Passengers_info: bookedticket.Passengers_info,
		}

		//result, err := bookedticketCollection.InsertOne(ctx, newBookedTicket)
		result, err := bookedticketrepo.Insert(newBookedTicket)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// iid := fmt.Sprintf("%v", result.InsertedID)
		// new_produce_ticket := kafka_booking_ticket{
		// 	insertedid:   iid,
		// 	bookedticket: newBookedTicket,
		// }
		err = userrepo.UpdateBookedTicket(bookedticket.User_id, bookedticket.Id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		// go produce_booked_ticket(new_produce_ticket)

		c.JSON(http.StatusCreated, responses.BookedTicketResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetBookedTicket() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		bookedticketId := c.Param("bookedticketid")
		var bookedticket models.BookedTicket
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(bookedticketId)

		//err := bookedticketCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&bookedticket)
		bookedticket, err := bookedticketrepo.Read(objId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BookedTicketResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//check and update avaiable tickets
		var availticket models.AvailTicket

		//err := availticketCollection.FindOne(ctx, bson.M{"train_id": bookedticket.Train_id}).Decode(&availticket)
		availticket, err = availticketrepo.ReadTrainId(bookedticket.Train_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BookedTicketResponse{Status: http.StatusInternalServerError, Message: "Incorrect train id", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if availticket.Capacity == 0 {
			c.JSON(http.StatusInternalServerError, responses.BookedTicketResponse{Status: http.StatusInternalServerError, Message: "No tickets available", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// use this using kafka in future
		go kafka.Produce_booked_ticket_for_avail(bookedticket.Train_id, true)
		//update := bson.M{"capacity": availticket.Capacity - 1}
		//, err = availticketCollection.UpdateOne(ctx, bson.M{"trainid": bookedticket.Train_id}, bson.M{"$set": update})
		availticket.Capacity += 1
		_, err = availticketrepo.Update(availticket, availticket.Id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AvailTicketResponse{Status: http.StatusInternalServerError, Message: "error in updating capacity", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.BookedTicketResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": bookedticket}})
	}
}

func DeleteBookedTicket() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		bookedticketId := c.Param("bookedticketid")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(bookedticketId)

		// result, err := bookedticketCollection.DeleteOne(ctx, bson.M{"_id": objId})

		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, responses.BookedTicketResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		// 	return
		// }

		// if result.DeletedCount < 1 {
		// 	c.JSON(http.StatusNotFound,
		// 		responses.BookedTicketResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "BookedTicket with specified ID not found!"}},
		// 	)
		// 	return
		// }

		bookedticket, err := bookedticketrepo.Read(objId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BookedTicketResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		resultcnt, err := bookedticketrepo.Delete(objId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BookedTicketResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		go kafka.Produce_booked_ticket_for_avail(bookedticket.Train_id, false)

		if resultcnt.(int) < 1 {
			c.JSON(http.StatusNotFound,
				responses.BookedTicketResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "BookedTicket with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.BookedTicketResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "BookedTicket successfully deleted!"}},
		)
	}
}

// func produce_booked_ticket(nbt kafka_booking_ticket) {
// 	l := log.New(os.Stdout, "kafka producer", 0)
// 	w := kafka.NewWriter(kafka.WriterConfig{
// 		Brokers: []string{brokerAddress},
// 		Topic:   topic,
// 		Logger:  l,
// 	})

// 	bytes, _ := json.Marshal(nbt.bookedticket)
// 	err := w.WriteMessages(context.Background(), kafka.Message{
// 		Key:   []byte(nbt.insertedid),
// 		Value: []byte(bytes),
// 	})
// 	if err != nil {
// 		panic("could not write message " + err.Error())
// 	}
// }

// func consume_booked_ticket() {
// 	l := log.New(os.Stdout, "kafka producer", 0)
// 	r := kafka.NewReader(kafka.ReaderConfig{
// 		Brokers: []string{brokerAddress},
// 		Topic:   topic,
// 		Logger:  l,
// 	})

// 	for {
// 		// the `ReadMessage` method blocks until we receive the next event
// 		msg, err := r.ReadMessage(context.Background())
// 		if err != nil {
// 			panic("could not read message " + err.Error())
// 		}
// 		// after receiving the message, log its value
// 		fmt.Println("received: ", string(msg.Value))
// 		nbtr := models.BookedTicket{}
// 		json.Unmarshal([]byte(msg.Value), &nbtr)
// 		fmt.Println(nbtr)
// 		update_tickets_booked := bson.M{"tickets_booked": string(msg.Key)}
// 		_, err = userCollection.UpdateOne(context.Background(), bson.M{"_id": nbtr.User_id}, bson.M{"$push": update_tickets_booked})
// 		if err != nil {
// 			panic("could not update booked ticket " + err.Error())
// 		}
// 	}
// }
