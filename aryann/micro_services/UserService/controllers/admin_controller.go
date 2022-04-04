package controllers

import (
	"UserService/models"
	"UserService/responses"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var avalidate = validator.New()

// GetAdmin godoc
// @Summary      Fetch An Admin
// @Description  Get Admin details by providing the adminid
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        adminid 		body	string  true  "admin unique id"
// @Success      200  {object}  models.Admin
// @Failure      400  {number} 	http.StatusBadRequest
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /admin/:adminid [get]
func GetAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		adminId := c.Param("adminid")
		var admin models.Admin
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(adminId)

		admin, err := adminRepo.Read(objId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.AdminResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": admin}})
	}
}

// DelAdmin godoc
// @Summary      Delete An Admin
// @Description  Delete an Admin by providing the adminid
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        adminid 		body	string  true  "admin unique id"
// @Success      200  {object}  models.Admin
// @Failure      400  {number} 	http.StatusBadRequest
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /admin/:adminid [delete]
func DeleteAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		adminId := c.Param("adminid")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(adminId)

		result, err := adminRepo.Delete(objId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		admindel, err := adminRepo.Delete(objId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK,
			responses.AdminResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Admin successfully deleted!", "result": result, "admin": admindel}},
		)
	}
}
