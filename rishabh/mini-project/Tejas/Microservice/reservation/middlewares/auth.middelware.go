package middlewares

import (
	"net/http"
	"reservationService/services"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthenticateJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Auth token not found"})
			c.Abort()
		}
		token := strings.Fields(authHeader)
		tokenString := token[1]
		verified, _ := services.ValidateToken(tokenString)
		if !verified {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid auth token"})
			c.Abort()
			return
		}
		claims, err := services.GetClaimsFromToken(tokenString)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid auth token"})
			c.Abort()
			return
		}
		c.Set("user_details", claims)
	}
}

func OnlyAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userDetails services.SignedDetails = c.MustGet("user_details").(services.SignedDetails)
		if !userDetails.IsAdmin {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Only admin can perform this action"})
			c.Abort()
			return
		} else {
			c.Next()
		}
	}
}
