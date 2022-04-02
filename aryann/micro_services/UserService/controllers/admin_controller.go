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

const layout = "Jan 2, 2006 at 3:04pm (MST)"

// func init() {
// 	go kafka.Consume_purchased_ticket()
// }

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
