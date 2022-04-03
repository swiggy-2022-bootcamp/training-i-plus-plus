package service

import (
	"Reservations/config"
	errors "Reservations/errors"
	"Reservations/kafka"
	"Reservations/middleware"
	model "Reservations/model"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var mongoURL string = config.MONGO_URL
var ticketCollection *mongo.Collection

func init() {
	client, _ = mongo.NewClient(options.Client().ApplyURI(mongoURL))
	ticketCollection = client.Database("TrainTicketLelo").Collection("Reservations")
}

func BuyTicket(body *io.ReadCloser) (result *mongo.InsertOneResult, err error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	client.Connect(ctx)

	var ticketBought model.Reservation
	json.NewDecoder(*body).Decode(&ticketBought)

	userId := ticketBought.UserId
	if !IsValidUser(userId) {
		return nil, errors.UserNotFoundError()
	}

	success, errorResponse, errorProductIndex := UpdateTicketCount(userId, ticketBought.TrainIDs, -1)
	if !success {
		errorMessage := ReadCloserToString(errorResponse.Body) + ". Ticket Id: " + ticketBought.TrainIDs[*errorProductIndex] + " (Order rolled back)"
		return nil, &errors.PurchaseError{Status: http.StatusBadRequest, ErrorMessage: errorMessage}
	}

	ticketBought.PurchaseDate = time.Now()
	ticketBought.DepartureDate = ticketBought.PurchaseDate.AddDate(0, 0, 2)
	ticketBought.Status = "Payment Pending"

	ctx, _ = context.WithTimeout(context.Background(), time.Second*10)
	result, _ = ticketCollection.InsertOne(ctx, ticketBought)
	ticketId := result.InsertedID.(primitive.ObjectID).Hex()

	ctx, _ = context.WithTimeout(context.Background(), time.Minute*10)
	kafka.Produce(ctx, nil, []byte("ticketId: "+ticketId+" --- status: "+ticketBought.Status))

	return
}

func GetTickets(userId string) (tickets []model.Reservation, err error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	client.Connect(ctx)

	if !IsValidUser(userId) {
		return nil, errors.UserNotFoundError()
	}

	ctx, _ = context.WithTimeout(context.Background(), time.Second*10)
	cursor, _ := ticketCollection.Find(ctx, bson.M{"userid": userId})

	for cursor.Next(ctx) {
		var ticket model.Reservation
		cursor.Decode(&ticket)
		tickets = append(tickets, ticket)
	}
	return
}

func TicketPayment(ticketId string) (successMessage *string, err error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	client.Connect(ctx)

	objectId, err := primitive.ObjectIDFromHex(ticketId)
	if err != nil {
		return nil, errors.MalformedIdError()
	}

	result := ticketCollection.FindOne(ctx, bson.M{"_id": objectId})

	if result.Err() != nil && result.Err() == mongo.ErrNoDocuments {
		return nil, errors.IdNotFoundError()
	}

	var ticket model.Reservation
	result.Decode(&ticket)

	if ticket.Status == "Payment Done" {
		return nil, errors.OrderAlreadyPaidForError()
	}

	ticket.Status = "Payment Done"

	_, error := ticketCollection.UpdateByID(ctx, objectId, bson.M{"$set": ticket})

	if error != nil {
		return nil, errors.InternalServerError()
	}

	ctx, _ = context.WithTimeout(context.Background(), time.Minute*10)
	kafka.Produce(ctx, nil, []byte("ticketId: "+ticketId+" --- status: "+ticket.Status))

	str := "Successfully Paid For The Ticket"
	successMessage = &str
	return
}

func IsValidUser(userId string) bool {
	jwtToken, _ := middleware.GenerateJWT(userId, model.Admin)
	url := "http://localhost:8001/users/"

	var bearer = "Bearer " + jwtToken

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode != http.StatusOK {
		return false
	}
	return true
}

func UpdateTicketCount(userId string, TrainIDs []string, quantity int) (success bool, errorResponse *http.Response, errorProductIndex *int) {
	jwtToken, _ := middleware.GenerateJWT(userId, model.Admin)

	var bearer = "Bearer " + jwtToken

	for index, trainId := range TrainIDs {
		url := "http://localhost:8003/trains/" + trainId + "/" + strconv.Itoa(quantity)

		req, _ := http.NewRequest("POST", url, nil)

		req.Header.Add("Authorization", bearer)

		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			log.Fatalln(err)
		}

		if resp.StatusCode != http.StatusOK {
			UpdateTicketCount(userId, TrainIDs[:index], 1)
			return false, resp, &index
		}
	}
	return true, nil, nil
}

func ReadCloserToString(body io.ReadCloser) (message string) {
	json.NewDecoder(body).Decode(&message)
	return
}
