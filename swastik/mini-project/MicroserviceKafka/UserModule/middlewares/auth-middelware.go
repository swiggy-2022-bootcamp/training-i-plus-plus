package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/swastiksahoo153/train-reservation-system/services"
	"net/http"
	"strings"
	"fmt"
)

func AuthenticateJWT() gin.HandlerFunc{
	return func(c *gin.Context){
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error":"Auth token not found"})
			c.Abort()
		}
		token := strings.Fields(authHeader)
		fmt.Println("Token: ", token)
		tokenString := token[0]
		fmt.Println(tokenString)
		verified,_ := services.ValidateToken(tokenString)
		if !verified {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"Invalid auth token"})
			c.Abort()
			return
		}else{
			c.Next()
		}
	}
}