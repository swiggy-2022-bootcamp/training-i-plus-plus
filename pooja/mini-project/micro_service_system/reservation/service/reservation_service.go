package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"reservation/database"
	"reservation/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var bookingCollection *mongo.Collection = database.GetCollection(database.DB, "bookings")

func BookTickets() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var bookingReq model.BookTicketsRequest
		//bookingCollection.DeleteMany(ctx, bson.M{})
		if err := c.BindJSON(&bookingReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error()})
			return
		}
		//check if the train exists and if seats are available in it
		success, errorResponse := CheckAndUpdateSeatAvailabilty(bookingReq.UserName, bookingReq.TrainNumber, bookingReq.NumberOfSeats, false)
		fmt.Println("success", success)
		if !success {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": errorResponse})
			return
		}
		model.GlobalPNR += 1
		bookingResp := model.Booking{
			ID:            primitive.NewObjectID(),
			UserName:      bookingReq.UserName,
			PNR:           model.GlobalPNR,
			NumberOfSeats: bookingReq.NumberOfSeats,
			TrainNumber:   bookingReq.TrainNumber,
			DepartureDate: bookingReq.DepartureDate,
			BookingStatus: "CONFIRMED",
		}
		result, insertErr := bookingCollection.InsertOne(ctx, bookingResp)
		if insertErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": insertErr})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "tickets booked", "data": result})
	}
}

func CancelBooking() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var cancelBookingReq model.CancelBookingRequest
		var booking model.Booking

		if err := c.BindJSON(&cancelBookingReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error()})
			return
		}

		//get the booking from the datastore

		if err := bookingCollection.FindOne(ctx, bson.M{"pnr": cancelBookingReq.PNR}).Decode(&booking); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusBadRequest, "message": "error in fetching booking", "error": err.Error()})
			return
		}
		//check if the booking is already cancelled
		if booking.BookingStatus == "CANCELLED" {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusBadRequest, "message": "booking is already cancelled"})
			return
		}
		//cancel the booking
		booking.BookingStatus = "CANCELLED"
		_, err := bookingCollection.UpdateOne(ctx, bson.M{"pnr": cancelBookingReq.PNR}, bson.M{"$set": booking})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": err.Error(), "message": "error while cancelling the booking"})
			return
		}
		//increment the available seat count
		success, errorResponse := CheckAndUpdateSeatAvailabilty(cancelBookingReq.UserName, booking.TrainNumber, booking.NumberOfSeats, true)
		fmt.Println("success", success)
		if !success {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": errorResponse})
			return
		}

		// 		//increase the count of available seats
		// 		var train model.Train
		// 		if err := trainCollection.FindOne(ctx, bson.M{"trainnumber": booking.TrainNumber, "departuredate": booking.DepartureDate}).Decode(&train); err != nil {
		// 			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": err.Error()})
		// 			return
		// 		}
		// 		train.AvailableSeatCount += booking.NumberOfSeats
		// 		fmt.Print(train)
		// 		_, err = trainCollection.UpdateOne(ctx, bson.M{"trainnumber": booking.TrainNumber, "departuredate": booking.DepartureDate}, bson.M{"$set": train})
		// 		if err != nil {
		// 			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": err.Error(), "message": "error while cancelling the booking"})
		// 			return
		// 		}

		// c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "tickets booked", "data": booking})
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
	}
}

func ViewBookings() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var getBookingsReq model.BookingsByUserRequest
		var bookings []model.Booking

		if err := c.BindJSON(&getBookingsReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error()})
			return
		}

		//get the bookings from the datastore
		results, err := bookingCollection.Find(ctx, bson.M{"username": getBookingsReq.UserName})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": err.Error()})
			return
		}
		fmt.Println(getBookingsReq.UserName, results)
		defer results.Close(ctx)
		for results.Next(ctx) {
			var booking model.Booking
			if err = results.Decode(&booking); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": err.Error()})
				return
			}
			singleBooking := model.Booking{
				ID:            booking.ID,
				UserName:      booking.UserName,
				PNR:           booking.PNR,
				NumberOfSeats: booking.NumberOfSeats,
				TrainNumber:   booking.TrainNumber,
				DepartureDate: booking.DepartureDate,
				BookingStatus: booking.BookingStatus,
			}
			bookings = append(bookings, singleBooking)
		}
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "success", "data": bookings})
	}
}

func GetAllBookings() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var bookings []model.Booking
		defer cancel()
		results, err := bookingCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": err.Error()})
			return
		}
		defer results.Close(ctx)
		fmt.Println(results)
		for results.Next(ctx) {
			var singleBooking model.Booking
			if err = results.Decode(&singleBooking); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": err.Error()})
				return
			}

			bookings = append(bookings, singleBooking)
		}
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "success", "data": bookings})
	}
}

type respAvailability struct {
	status  string         `json:"status"`
	errResp *http.Response `json:"error"`
}

func CheckAndUpdateSeatAvailabilty(username string, train_number string, number_of_seats int, incrementCount bool) (sucess bool, err error) {
	//jwtToken, _ := middleware.GenerateJWT(userId, mockdata.Admin)

	// Create a Bearer string by appending string access token
	//var bearer = "Bearer " + jwtToken
	apiUrl := "http://localhost:6002/checkavailability"
	fmt.Println(train_number, " ", number_of_seats)
	// Create a new request using http
	req, _ := http.NewRequest("GET", apiUrl, nil)
	q := req.URL.Query()
	q.Add("trainnumber", train_number)
	q.Add("numofseats", strconv.Itoa(number_of_seats))
	q.Add("incrementcount", strconv.FormatBool(incrementCount))
	req.URL.RawQuery = q.Encode()
	fmt.Print(req.URL.String())
	req.Close = true
	// add authorization header to the req
	//req.Header.Add("token", bearer)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	fmt.Println("resp", resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		return false, errors.New("cannot update train availablity")
	}
	//bodyBytes, _ := ioutil.ReadAll((resp.Body))
	//fmt.Println(bodyBytes)
	// fmt.Println(string(bodyBytes), "--------", resp.StatusCode)
	// var respObj respAvailability
	// json.Unmarshal(bodyBytes, &respObj)
	// fmt.Println("response object", respObj)
	// if !respObj.success || respObj.errResp != nil {
	// 	return false, respObj.errResp
	// }
	return true, nil
}
