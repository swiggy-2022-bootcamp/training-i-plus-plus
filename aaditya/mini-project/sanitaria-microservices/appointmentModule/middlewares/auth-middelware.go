package middlewares

import (
	"github.com/gin-gonic/gin"
	"sanitaria-microservices/doctorModule/services"
	"net/http"
	"strings"
)

func AuthenticateJWT() gin.HandlerFunc{
	return func(c *gin.Context){
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error":"Auth token not found"})
			c.Abort()
		}
		token := strings.Fields(authHeader)
		tokenString := token[1]
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