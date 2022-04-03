package middleware

import (
	"Reservations/kafka"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IfAuthorized(endPoint func(c *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {

		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")

		if len(authHeader) == 0 {
			c.JSON(http.StatusUnauthorized, "Unauthorized API Call")
			return
		}

		tokenString := authHeader[len(BEARER_SCHEMA):]
		fmt.Println(tokenString)

		userId, userRole, error := ValidateToken(tokenString)

		if error != nil {
			ctx := context.Context(context.Background())
			kafka.Produce(ctx, nil, []byte("API Access Not Given To Unauthorized User"))
			c.JSON(http.StatusUnauthorized, "Invalid token. "+error.Error())
			return
		}

		c.Params = append(c.Params, gin.Param{
			Key:   "acessorUserId",
			Value: userId,
		},
			gin.Param{
				Key:   "acessorUserRole",
				Value: strconv.Itoa(int(userRole)),
			},
		)

		endPoint(c)
	}
}
