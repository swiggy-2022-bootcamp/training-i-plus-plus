package middleware

import (
	"User-Service/kafka"
	"context"
	"fmt"
	"net/http"

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

		jwtToken, error := ValidateToken(tokenString)

		if error != nil {
			ctx := context.Context(context.Background())
			kafka.Produce(ctx, nil, []byte("Unauthorized API access averted"))
			c.JSON(http.StatusUnauthorized, "Invalid token. "+error.Error())
		}

		if jwtToken.Valid {
			endPoint(c)
		}
	}
}
