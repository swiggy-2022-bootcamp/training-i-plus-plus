package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetMyAppointments godoc
// @Summary get Appointments of logged in user
// @Tags Appointment
// @Success 200
// @Failure 500
// @Router /appointments [get]
// @Security ApiKeyAuth
func (cont Controller) GetMyAppointments(c *gin.Context) {

	userId := c.GetString("userId")

	response, err := http.Get("http://localhost:7451/appointment/user/" + userId)

	if err != nil {
		log.Println("Error while fetching appointments " + err.Error())
		c.AbortWithStatusJSON(500, gin.H{
			"error": "Error while fetching appointments from doctor",
		})
		return
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	var res interface{}
	if err = json.Unmarshal(responseData, &res); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, res)
}
