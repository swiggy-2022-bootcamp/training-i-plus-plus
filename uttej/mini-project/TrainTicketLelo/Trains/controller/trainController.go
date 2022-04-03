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

func CreateTrain(c *gin.Context) {
	acessorUserRole, _ := strconv.Atoi(c.Param("acessorUserRole"))

	if models.Role(acessorUserRole) != models.Admin {
		c.JSON(http.StatusUnauthorized, errors.AccessDenied())
		return
	}

	result := service.CreateTrain(&c.Request.Body)
	c.JSON(http.StatusOK, result)
}

func GetTrains(c *gin.Context) {
	allTrains := service.GetTrains()
	c.JSON(http.StatusOK, allTrains)
}

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
