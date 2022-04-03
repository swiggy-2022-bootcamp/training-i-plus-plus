package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"
	"userService/kafka"
	log "userService/logger"
	"userService/models"
	pb "userService/protobuf"
	"userService/repository"
	"userService/responses"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
)

//var userCollection *mongo.Collection = config.GetCollection(config.DB, "users")
//var bookedticketCollection *mongo.Collection = config.GetCollection(config.DB, "bookedtickets")

var (
	errLog           = log.ErrorLogger.Println
	infoLog          = log.ErrorLogger.Println
	avalidate        = validator.New()
	address          = "localhost:6010"
	userrepo         repository.UserRepository
	bookedticketrepo repository.BookedTicketRepository
	trainrepo        repository.TrainRepository
	c                pb.AuthenticationServiceClient
	c1               pb.AvailTicketServiceClient
	c2               pb.ChargeServiceClient
)

func init() {
	go kafka.Consume_avail_ticket()
	go kafka.Consume_train()

	//configure grpc client
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		errLog("Error while making connection, %v", err)
		fmt.Printf("Error while making connection, %v\n", err)
	}

	// Create a client instance
	c = pb.NewAuthenticationServiceClient(conn)
}

func CheckAuthorized(group string) gin.HandlerFunc {
	return func(co *gin.Context) {
		bearToken := co.Request.Header.Get("Authorization")
		//normally Authorization the_token_xxx
		strArr := strings.Split(bearToken, " ")
		if len(strArr) != 2 {
			respondWithError(co, 401, "No bearer token")
			return
		}
		resp, err := c.Authenticate(
			context.Background(),
			&pb.AuthenticateRequest{
				Group: group,
				Token: strArr[1],
			})
		if err != nil || !resp.Confirmation {
			respondWithError(co, 401, err)
			return
		}
		infoLog("Autoriztion " + resp.Message)
		co.Next()
	}
}

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

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
		if validationErr := avalidate.Struct(&user); validationErr != nil {
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
		if validationErr := avalidate.Struct(&user); validationErr != nil {
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

func checkAvailticketgrpc(train_id string, capacity int) (*pb.AvailTicketResponse, error) {
	resp, err := c1.GetTicketConfirmation(context.Background(), &pb.AvailTicketRequest{
		TrainId:      train_id,
		NumOfTickets: uint32(capacity),
	})
	return resp, err
}

func payBookedticketgrpc(amount int, email string, trainid string) (*pb.ChargeResponse, error) {
	resp, err := c2.Charge(context.Background(), &pb.ChargeRequest{
		Amount:       uint32(amount),
		Receiptemail: email,
		Ticketid:     trainid,
	})
	return resp, err
}

func CreateBookedTicket() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var bookedticket models.BookingTicket
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&bookedticket); err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error in binding", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := avalidate.Struct(&bookedticket); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error in validating", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		//----------------------- GRPC
		// var availticket models.AvailTicket

		// //err := availticketCollection.FindOne(ctx, bson.M{"train_id": bookedticket.Train_id}).Decode(&availticket)
		// availticket, err := availticketrepo.ReadTrainId(bookedticket.Train_id)
		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, responses.BookedTicketResponse{Status: http.StatusInternalServerError, Message: "Incorrect train id", Data: map[string]interface{}{"data": err.Error()}})
		// 	return
		// }

		// if availticket.Capacity == 0 {
		// 	c.JSON(http.StatusBadRequest, responses.BookedTicketResponse{Status: http.StatusInternalServerError, Message: "No tickets available", Data: map[string]interface{}{"data": err.Error()}})
		// 	return
		// }

		// // --------------------- kafka

		// //update := bson.M{"capacity": availticket.Capacity - 1}
		// //, err = availticketCollection.UpdateOne(ctx, bson.M{"trainid": bookedticket.Train_id}, bson.M{"$set": update})
		// availticket.Capacity -= 1
		// _, err = availticketrepo.Update(availticket, availticket.Id)
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, responses.AvailTicketResponse{Status: http.StatusInternalServerError, Message: "error in updating capacity", Data: map[string]interface{}{"data": err.Error()}})
		// 	return
		// }

		// var trainbooked models.Train

		// //----------------------- GRPC
		// //err = trainCollection.FindOne(ctx, bson.M{"_id": bookedticket.Train_id}).Decode(&trainbooked)
		// trainbooked, err = trainrepo.Read(bookedticket.Train_id)
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, responses.BookedTicketResponse{Status: http.StatusInternalServerError, Message: "error in train find", Data: map[string]interface{}{"data": err.Error()}})
		// 	return
		// }

		resp, err := checkAvailticketgrpc(bookedticket.Train_id.String(), len(bookedticket.Passengers_info))

		if err != nil || resp.Message != 0 {
			c.JSON(http.StatusBadRequest, responses.BookedTicketResponse{Status: http.StatusInternalServerError, Message: "Incorrect train id", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// grpc for payment if succeds move forward

		usr, err := userrepo.Read(bookedticket.User_id)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.BookedTicketResponse{Status: http.StatusInternalServerError, Message: "Incorrect user id", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		resp1, err := payBookedticketgrpc(bookedticket.Amount_paid, bookedticket.Train_id.String(), usr.Email)

		if err != nil {
			c.JSON(http.StatusBadRequest, responses.BookedTicketResponse{Status: http.StatusInternalServerError, Message: "Payment Unsuccesfull", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		infoLog(resp1.Message)
		go kafka.Produce_booked_ticket_for_avail(bookedticket.Train_id, true)

		//--------------------------
		newBookedTicket := models.BookedTicket{
			Train_id:        bookedticket.Train_id,
			User_id:         bookedticket.User_id,
			Departure:       resp.Station1,
			Arrival:         resp.Station2,
			Departure_time:  resp.DepartureTime,
			Arrival_time:    resp.ArrivalTime,
			Passengers_info: bookedticket.Passengers_info,
		}

		//result, err := bookedticketCollection.InsertOne(ctx, newBookedTicket)
		result, err := bookedticketrepo.Insert(newBookedTicket)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// iid := fmt.Sprintf("%v", result.InsertedID)
		// new_produce_ticket := kafka_booking_ticket{
		// 	insertedid:   iid,
		// 	bookedticket: newBookedTicket,
		// }
		err = userrepo.UpdateBookedTicket(bookedticket.User_id, bookedticket.Id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
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

func GetAllTrains() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var trains []models.Train
		defer cancel()

		trains, err := trainrepo.ReadAll()

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BookedTicketResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK,
			responses.BookedTicketResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": trains}},
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
