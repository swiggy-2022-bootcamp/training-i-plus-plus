package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"ticket_service/responses"

	"github.com/gin-gonic/gin"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	mySigningKey = []byte("fdsgdgjhgfj")
)

func GetKey() []byte {
	return mySigningKey
}
func Authenticate(role string, checkRole bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		bearToken := c.Request.Header.Get("Authorization")

		strArr := strings.Split(bearToken, " ")
		if len(strArr) != 2 {
			c.AbortWithStatusJSON(http.StatusBadRequest, responses.TicketResponse{Status: http.StatusInternalServerError, Message: "No token present"})
			return
		}

		token, err := jwt.Parse(strArr[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf(("invalid token"))
			}

			return GetKey(), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, responses.TicketResponse{Status: http.StatusInternalServerError, Message: "Internal Error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusBadRequest, responses.TicketResponse{Status: http.StatusInternalServerError, Message: "Invalid token"})
			return
		}

		if checkRole && token.Claims.(jwt.MapClaims)["role"] != role {
			c.AbortWithStatusJSON(http.StatusBadRequest, responses.TicketResponse{Status: http.StatusInternalServerError, Message: "Unauthorized User"})
			return
		}
		c.Set("user_details", token.Claims)
		c.Next()
	}
}