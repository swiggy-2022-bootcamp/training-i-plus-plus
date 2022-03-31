package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-kafka-microservice/ListingService/services"
)

type ListingMiddleware struct {
	ListingService services.ListingService
}

func NewListingMiddleware(listingService services.ListingService) *ListingMiddleware {
	return &ListingMiddleware{
		ListingService: listingService,
	}
}

func (im *ListingMiddleware) AuthorizeUser() gin.HandlerFunc {
	return func(gctx *gin.Context) {
		authHeaderReq := gctx.Request.Header.Get("Authorization")
		authHeader := strings.Split(authHeaderReq, "Bearer ")
		if len(authHeader) < 2 {
			gctx.JSON(http.StatusBadRequest, gin.H{"message": "Provide valid token."})
			gctx.Abort()
			return
		}

		fmt.Println(authHeader[1])
		// TODO: Need to use refreshed token
		_, err := im.ListingService.AuthorizeUser(authHeader[1])
		if err != nil {
			gctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			gctx.Abort()
			return
		}
		gctx.Next()
	}
}
