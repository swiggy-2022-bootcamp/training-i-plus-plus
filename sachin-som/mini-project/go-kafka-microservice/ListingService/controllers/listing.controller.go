package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-kafka-microservice/ListingService/services"
)

type ListingController struct {
	ListingService services.ListingService
}

func NewListingController(listingService services.ListingService) *ListingController {
	return &ListingController{
		ListingService: listingService,
	}
}

func (lc *ListingController) ShowProducts(gctx *gin.Context) {
	gctx.JSON(http.StatusOK, nil)
}
