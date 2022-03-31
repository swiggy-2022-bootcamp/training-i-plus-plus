package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-kafka-microservice/InventoryService/services"
)

type InventoryMiddleware struct {
	InventoryService services.InventoryServices
}

func NewInventoryMiddleware(inventoryService services.InventoryServices) *InventoryMiddleware {
	return &InventoryMiddleware{
		InventoryService: inventoryService,
	}
}

func (im *InventoryMiddleware) AuthorizeUser() gin.HandlerFunc {
	return func(gctx *gin.Context) {
		authHeaderReq := gctx.Request.Header.Get("Authorization")
		authHeader := strings.Split(authHeaderReq, "Bearer ")
		if len(authHeader) < 2 {
			gctx.JSON(http.StatusBadRequest, gin.H{"message": "Provide valid token."})
			gctx.Abort()
			return
		}

		fmt.Println(authHeader[1])
		token, err := im.InventoryService.AuthorizeUser(authHeader[1])
		if err != nil {
			gctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			gctx.Abort()
			return
		}
		fmt.Println(token)
		gctx.Next()
	}
}
