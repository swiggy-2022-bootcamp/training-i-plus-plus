package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-kafka-microservice/ListingService/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	products, err := lc.ListingService.ShowProducts()
	if err != nil {
		gctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	gctx.JSON(http.StatusOK, products)
}

func (lc *ListingController) MakeOrder(gctx *gin.Context) {
	ownerId, err := primitive.ObjectIDFromHex(gctx.Param("userId"))
	if err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	productId, err := primitive.ObjectIDFromHex(gctx.Param("productId"))
	if err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := lc.ListingService.MakeOrder(productId, ownerId); err != nil {
		gctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	gctx.JSON(http.StatusOK, gin.H{"message": "Your Order is initiated successfully, you will notify further once the order got placed successfully."})
}
