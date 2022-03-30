package middleware

import (
	"net/http"
	"strings"
	JWTManager "swiggy/gin/lib/helpers"

	"github.com/gin-gonic/gin"
)

func CheckAuthMiddleware(c *gin.Context) {
	token := c.Request.Header["Authorization"]
	if len(token) > 0 {
		ok, err := JWTManager.Manager.Verify(strings.Split(token[0], "Bearer ")[1])
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		} else {
			c.Set("User", ok.ID)
			c.Set("Role", ok.Role)
			c.Next()
		}
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Auth Token Not supplied"})
		return
	}
}

func CheckAdminRole(c *gin.Context) {
	if c.GetString("Role") == "admin" {
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Role not authorized"})
		return
	}
}
