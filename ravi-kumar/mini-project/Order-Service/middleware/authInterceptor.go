package middleware

import (
	"Order-Service/kafka"
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

		//authHeader will be "Bearer <bearer token>", so extract just the <bearer token> from this.
		tokenString := authHeader[len(BEARER_SCHEMA):]
		fmt.Println(tokenString)

		userId, userRole, error := ValidateToken(tokenString)

		if error != nil {
			ctx := context.Context(context.Background())
			kafka.Produce(ctx, nil, []byte("Unauthorized API access averted"))
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
