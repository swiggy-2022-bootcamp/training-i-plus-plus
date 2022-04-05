package middleware

import (
	"net/http"
	"userService/configs"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	
)

var mySigningKey = []byte(configs.EnvSecretKeyJWT())

func IsUserAuthorized(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := c.Request.Header.Get("Authorization")

		if authToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Auth token not found"})
			return
		}

		token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid auth token"})
			}
			return mySigningKey, nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotImplemented, gin.H{"error": "Invalid signin method"})
			return
		}
		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Invalid auth token"})
			return
		}
		for _ , role:= range roles {
			if token.Claims.(jwt.MapClaims)["role"] == role {
					c.Next()
					return;
				}	
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Inavalid user token"})
		return

	}
}
