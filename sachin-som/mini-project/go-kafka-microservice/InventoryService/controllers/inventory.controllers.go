package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-kafka-microservice/InventoryService/models"
	"github.com/go-kafka-microservice/InventoryService/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InventoryControllers struct {
	InventoryService services.InventoryServices
}

func NewInventoryControllers(inventoryService services.InventoryServices) *InventoryControllers {
	return &InventoryControllers{
		InventoryService: inventoryService,
	}
}

func (ic *InventoryControllers) RegisterInventory(gctx *gin.Context) {
	var inventory models.Inventory
	if err := gctx.ShouldBindJSON(&inventory); err != nil {
		gctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	inventoryId, err := ic.InventoryService.RegisterInventory(&inventory)
	if err != nil {
		gctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	gctx.JSON(http.StatusCreated, gin.H{"id": inventoryId})
}

func (ic *InventoryControllers) AddProduct(gctx *gin.Context) {
	inventoryId, _ := primitive.ObjectIDFromHex(gctx.Param("id"))
	var prouduct models.Product
	if err := gctx.ShouldBindJSON(&prouduct); err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := ic.InventoryService.AddProduct(inventoryId, &prouduct); err != nil {
		gctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	gctx.JSON(http.StatusCreated, gin.H{"message": "Product Added to Inventory."})
}

func (ic *InventoryControllers) GetProduct(gctx *gin.Context) {
	inventoryId, _ := primitive.ObjectIDFromHex(gctx.Param("inventoryId"))
	productId, _ := primitive.ObjectIDFromHex(gctx.Param("prouductId"))

	product, err := ic.InventoryService.GetProduct(inventoryId, productId)
	if err != nil {
		gctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	gctx.JSON(http.StatusOK, gin.H{"message": product})
}

func (ic *InventoryControllers) RegisterInventoryRoutes(rg *gin.RouterGroup) {
	inventoryRouter := rg.Group("/")
	inventoryRouter.POST("/register", ic.RegisterInventory)
	inventoryRouter.POST("/:id/product/add", ic.AddProduct)
	inventoryRouter.GET("/:inventoryId/product/:productId", ic.GetProduct)
}
