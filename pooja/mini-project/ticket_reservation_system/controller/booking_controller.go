package controller

import (
	"context"
	"fmt"
	"net/http"
	"ticket_reservation_system/config"
	"ticket_reservation_system/helper"
	"ticket_reservation_system/model"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var bookingCollection *mongo.Collection = config.GetCollection(config.DB, "bookings")

func SearchTrains() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var searchReq model.SearchTrainRequest
		var trains []model.SearchTrainResponse

		if err := c.BindJSON(&searchReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error()})
			return
		}
		token := c.Request.Header.Get("Cookie")
		claims, msg := helper.ValidateToken(token)
		fmt.Println(claims, msg)
		/*var train model.Train
		if err := trainCollection.FindOne(ctx, bson.M{"departurestation": searchReq.DepartureStation}).Decode(&train); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": err.Error()})
			return
		}*/

		results, err := trainCollection.Find(ctx, bson.M{"departurestation": searchReq.DepartureStation,
			"arrivalstation": searchReq.ArrivalStation})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": err.Error()})
			return
		}
		defer results.Close(ctx)
		fmt.Println(searchReq.DepartureStation, searchReq.ArrivalStation, results)
		for results.Next(ctx) {
			var train model.Train
			if err = results.Decode(&train); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": err.Error()})
				return
			}
			if train.AvailableSeatCount > 0 {
				singleSearchResp := model.SearchTrainResponse{
					TrainNumber:        train.TrainNumber,
					TrainName:          train.TrainName,
					AvailableSeatCount: train.AvailableSeatCount,
					Fare:               train.Fare,
				}
				trains = append(trains, singleSearchResp)
			}
		}
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "success", "data": trains})
	}
}

func BookTickets() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var bookingReq model.BookTicketsRequest
		var train model.Train
		//bookingCollection.DeleteMany(ctx, bson.M{})
		if err := c.BindJSON(&bookingReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error()})
			return
		}
		if err := trainCollection.FindOne(ctx, bson.M{"trainnumber": bookingReq.TrainNumber}).Decode(&train); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": err.Error()})
			return
		}
		if train.AvailableSeatCount < bookingReq.NumberOfSeats {
			c.JSON(http.StatusBadGateway, gin.H{"error": "tickets are not available"})
			return
		}
		train.AvailableSeatCount -= bookingReq.NumberOfSeats
		_, err := trainCollection.UpdateOne(ctx, bson.M{"trainnumber": bookingReq.TrainNumber}, bson.M{"$set": train})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": err.Error()})
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
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error()})
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

		//increase the count of available seats
		var train model.Train
		if err := trainCollection.FindOne(ctx, bson.M{"trainnumber": booking.TrainNumber, "departuredate": booking.DepartureDate}).Decode(&train); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": err.Error()})
			return
		}
		train.AvailableSeatCount += booking.NumberOfSeats
		fmt.Print(train)
		_, err = trainCollection.UpdateOne(ctx, bson.M{"trainnumber": booking.TrainNumber, "departuredate": booking.DepartureDate}, bson.M{"$set": train})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": err.Error(), "message": "error while cancelling the booking"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "tickets booked", "data": booking})
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
