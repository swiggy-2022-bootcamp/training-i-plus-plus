package controllers

import (
	"PurchaseService/kafka"
	"PurchaseService/middlewares"
	"PurchaseService/models"
	"PurchaseService/repository"
	"PurchaseService/responses"
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validate = validator.New()
var purchasedRepo repository.PurchasedRepository

// PurchaseTicket godoc
// @Summary      Purchase A Ticket
// @Description  Purchase Ticket by providing ticket details
// @Tags         user
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Purchased
// @Failure      400  {number} 	http.StatusBadRequest
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /purchase [post]
func PurchaseTicket() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var purchased models.Purchased
		defer cancel()

		if err := c.BindJSON(&purchased); err != nil {
			c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error in binding", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if validationErr := validate.Struct(&purchased); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error in validating", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newpurchased := models.Purchased{
			Train_id:       purchased.Train_id,
			User_id:        purchased.User_id,
			Departure:      purchased.Departure,
			Arrival:        purchased.Arrival,
			Departure_time: purchased.Departure_time,
			Arrival_time:   purchased.Arrival_time,
			Cost:           purchased.Cost,
		}

		result, err := purchasedRepo.Create(newpurchased)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		go kafka.Produce_purchased_ticket(purchased.Train_id)

		c.JSON(http.StatusCreated, responses.PurchasedResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": "Ticket successfully purchased!", "purchased": result}})
	}
}

// GetPurchase godoc
// @Summary      Fetch a Purchased Ticket
// @Description  Get Purchased Ticket by providing ticket id
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        purchasedid 		body	string  true  "purchase unique id"
// @Success      200  {object}  models.Purchased
// @Failure      400  {number} 	http.StatusBadRequest
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /purchase/:purchasedid [get]
func GetPurchased() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		purchasedId := c.Param("purchasedid")
		var purchased models.Purchased
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(purchasedId)

		purchased, err := purchasedRepo.Read(objId)
		// err := purchasedCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&purchased)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PurchasedResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.PurchasedResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": purchased}})
	}
}

// DeletePurchase godoc
// @Summary      Delete a Purchased Ticket
// @Description  Delete Purchased Ticket by providing ticket id
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        purchasedid 		body	string  true  "purchase unique id"
// @Success      200  {object}  models.Purchased
// @Failure      400  {number} 	http.StatusBadRequest
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /purchase/:purchasedid [delete]
func DeletePurchased() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		purchasedId := c.Param("purchasedid")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(purchasedId)

		purchased, err := purchasedRepo.Read(objId)
		// result, err := purchasedCollection.DeleteOne(ctx, bson.M{"_id": objId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PurchasedResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		result, err := purchasedRepo.Delete(objId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PurchasedResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// if result.(int) < 1 {
		// 	c.JSON(http.StatusNotFound,
		// 		responses.PurchasedResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Purchased with specified ID not found!"}},
		// 	)
		// 	return
		// }

		c.JSON(http.StatusOK,
			responses.PurchasedResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Purchased successfully deleted!", "result": result, "purchased": purchased}},
		)
	}
}

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func IsAuthorized(group string) gin.HandlerFunc {
	return func(c *gin.Context) {
		bearToken := c.Request.Header.Get("Authorization")

		strArr := strings.Split(bearToken, " ")
		if len(strArr) != 2 {
			respondWithError(c, 401, "No bearer token")
			return
		}

		token, err := jwt.Parse(strArr[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				log.Fatal("Unexpected signing method")
				return nil, fmt.Errorf(("invalid signing method"))
			}

			return middlewares.GetMySigingKey(), nil
		})
		if err != nil {
			respondWithError(c, 501, err.Error())
			return
		}
		if !token.Valid {
			respondWithError(c, 401, "Invalid token")
			return
		}

		if token.Claims.(jwt.MapClaims)["group"] != group {
			respondWithError(c, 401, "unauthorized user")
			return
		}

		c.Next()
	}
}
