package controllers

import (
	errors "Trains/errors"
	models "Trains/models"
	service "Trains/service"
	"strconv"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateTrain godoc
// @Summary      Create a Train
// @Description  Create a train by providing the necessary details.
// @Tags         Train
// @Accept       json
// @Produce      json
// @Param		Train	body	models.Train	true	"id will be populated Automatically"
// @Success      200  {string}  responseBody
// @Failure      400  {number} 	http.StatusBadRequest
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /trains [post]
func CreateTrain(c *gin.Context) {
	acessorUserRole, _ := strconv.Atoi(c.Param("acessorUserRole"))

	if models.Role(acessorUserRole) != models.Admin {
		c.JSON(http.StatusUnauthorized, errors.AccessDenied())
		return
	}

	result := service.CreateTrain(&c.Request.Body)
	c.JSON(http.StatusOK, result)
}

// GetTrains godoc
// @Summary      Fetch All Trains
// @Description  Get All Trains & the details
// @Tags         Train
// @Accept       json
// @Produce      json
// @Success      200  {array}  models.Train
// @Failure      400  {number} 	http.StatusBadRequest
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /trains [get]
func GetTrains(c *gin.Context) {
	allTrains := service.GetTrains()
	c.JSON(http.StatusOK, allTrains)
}

// UpdateTicketCount godoc
// @Summary      Update the count of tickets for a train
// @Description  Update the tickets for a train by providing train id and ticket count. Count can be a positive or negative integer
// @Tags         Train
// @Accept       json
// @Produce      json
// @Param        TrainId 		body	string  true  "unique train id"
// @Param        TicketCount 		body	integer  true  "count of tickets"
// @Success      200  {string}  ticketCount
// @Failure      400  {number} 	http.StatusBadRequest
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /trains/:trainId/:updateCount [post]
func UpdateTicketCount(c *gin.Context) {
	acessorUserRole, _ := strconv.Atoi(c.Param("acessorUserRole"))
	if !(models.Role(acessorUserRole) == models.Admin) {
		c.JSON(http.StatusUnauthorized, errors.AccessDenied())
		return
	}

	var trainId string = c.Param("trainId")
	updateCount, err := strconv.Atoi(c.Param("updateCount"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Count Should Be An Integer")
		return
	}

	quantityAfterUpdation, error := service.UpdateTicketCount(trainId, updateCount)

	if error != nil {
		trainError, ok := error.(*errors.TrainError)
		if ok {
			c.JSON(trainError.Status, trainError.ErrorMessage)
			return
		} else {
			fmt.Println("Couldn't Update Count")
			return
		}
	}
	c.JSON(http.StatusOK, quantityAfterUpdation)
}

// GetTrain godoc
// @Summary      Fetch A Train
// @Description  Get Train details by providing the trainid
// @Tags         Train
// @Accept       json
// @Produce      json
// @Param        TrainId 		body	string  true  "unique train id"
// @Success      200  {object}  models.Train
// @Failure      400  {number} 	http.StatusBadRequest
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /trains/:trainId [get]
func GetTrainById(c *gin.Context) {
	var trainId string = c.Param("trainId")
	trainRetrieved, error := service.GetTrainById(trainId)

	if error != nil {
		trainError, ok := error.(*errors.TrainError)
		if ok {
			c.JSON(trainError.Status, trainError.ErrorMessage)
			return
		} else {
			fmt.Println("Couldn't Fetch The Train")
			return
		}
	}
	c.JSON(http.StatusOK, trainRetrieved)
}

// UpdateTrain godoc
// @Summary      Update A Train
// @Description  Update Train details by providing the trainid
// @Tags         Train
// @Accept       json
// @Produce      json
// @Param        TrainId 		body	string  true  "unique train id"
// @Success      200  {object}  models.Train
// @Failure      400  {number} 	http.StatusBadRequest
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /trains/:trainId [put]
func UpdateTrainById(c *gin.Context) {
	acessorUserRole, _ := strconv.Atoi(c.Param("acessorUserRole"))

	if models.Role(acessorUserRole) != models.Admin {
		c.JSON(http.StatusUnauthorized, errors.AccessDenied())
		return
	}

	var trainId string = c.Param("trainId")
	trainRetrieved, error := service.UpdateTrainById(trainId, &c.Request.Body)

	if error != nil {
		trainError, ok := error.(*errors.TrainError)
		if ok {
			c.JSON(trainError.Status, trainError.ErrorMessage)
			return
		} else {
			fmt.Println("Couldn't Update Train")
			return
		}
	}
	c.JSON(http.StatusOK, trainRetrieved)
}

// DeleteTrain godoc
// @Summary      Delete A Train
// @Description  Delete a Train by providing the trainid
// @Tags         Train
// @Accept       json
// @Produce      json
// @Param        TrainId 		body	string  true  "unique train id"
// @Success      200  {string}  successMessage
// @Failure      400  {number} 	http.StatusBadRequest
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /trains/:trainId [delete]
func DeleteTrainbyId(c *gin.Context) {
	acessorUserRole, _ := strconv.Atoi(c.Param("acessorUserRole"))

	if models.Role(acessorUserRole) != models.Admin {
		c.JSON(http.StatusUnauthorized, errors.AccessDenied())
		return
	}

	var trainId string = c.Param("trainId")
	successMessage, error := service.DeleteTrainbyId(trainId)

	if error != nil {
		trainError, ok := error.(*errors.TrainError)
		if ok {
			c.JSON(trainError.Status, trainError.ErrorMessage)
			return
		} else {
			fmt.Println("Couldn't Delete Train")
			return
		}
	}
	c.JSON(http.StatusOK, *successMessage)
}
