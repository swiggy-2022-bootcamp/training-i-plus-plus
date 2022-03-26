package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kafka-microservice/ListingService/controllers"
)

type ListingRoutes struct {
	ListingController *controllers.ListingController
}

func NewListingRoutes(listingController *controllers.ListingController) *ListingRoutes {
	return &ListingRoutes{
		ListingController: listingController,
	}
}

func (lr *ListingRoutes) RegisterListingRoutes(rg *gin.RouterGroup) {
	listingRouter := rg.Group("/")
	listingRouter.GET("/show/all", lr.ListingController.ShowProducts)
}
