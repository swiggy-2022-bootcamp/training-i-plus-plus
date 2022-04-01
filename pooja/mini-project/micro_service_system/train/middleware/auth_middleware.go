package middleware

import (
	"fmt"
	"net/http"
	"train/helper"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func AuthenticateJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := c.Request.Header.Get("token")
		if authToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "token not found"})
			c.Abort()
		}
		fmt.Print(authToken)
		if verified, err := helper.ValidateToken(authToken); !verified || err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid token"})
			log.Error("Invalid token")
			c.Abort()
			return
		}
		claims, err := helper.GetClaimsFromToken(authToken)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		c.Set("user_info", claims)
	}
}
