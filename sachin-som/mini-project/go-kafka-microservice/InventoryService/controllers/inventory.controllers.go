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

// RegisterInventory godoc
// @Summary      Creteas New Inventory
// @Description  inventory registration API
// @Tags         Inventory
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        inventory  body      requests.InventoryRequest  true  "inventory request structure"
// @Success      200        {object}  responses.IDResponse
// @Failure      400        {object}  responses.MessageResponse
// @Failure      500        {object}  responses.MessageResponse
// @Failure      502        {object}  responses.MessageResponse
// @Router       /register [post]
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

// AddProduct godoc
// @Summary      Adds New Product to A Inventory
// @Description  add product to inventory API
// @Tags         Product
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        product      body      requests.ProductRequest  true  "product request structure"
// @Param        inventoryId  path      string                   true  "unique Inventory Id"
// @Success      200          {object}  responses.MessageResponse
// @Failure      400          {object}  responses.MessageResponse
// @Failure      500          {object}  responses.MessageResponse
// @Failure      502          {object}  responses.MessageResponse
// @Router       /{inventoryId}/product/add [post]
func (ic *InventoryControllers) AddProduct(gctx *gin.Context) {
	inventoryId, _ := primitive.ObjectIDFromHex(gctx.Param("inventoryId"))
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

// GetProduct godoc
// @Summary      Get Product From An Inventory
// @Description  get product from inventory API
// @Tags         Product
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        inventoryId  path      string  true  "Unique Inventory Id"
// @Param        productId    path      string  true  "Unique Product Id"
// @Success      200          {object}  responses.ProductResponse
// @Failure      400          {object}  responses.MessageResponse
// @Failure      500          {object}  responses.MessageResponse
// @Failure      502          {object}  responses.MessageResponse
// @Router       /:inventoryId/product/:productId [get]
func (ic *InventoryControllers) GetProduct(gctx *gin.Context) {
	inventoryId, _ := primitive.ObjectIDFromHex(gctx.Param("inventoryId"))
	productId, _ := primitive.ObjectIDFromHex(gctx.Param("productId"))

	product, err := ic.InventoryService.GetProduct(inventoryId, productId)
	if err != nil {
		gctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	gctx.JSON(http.StatusOK, gin.H{"message": product})
}

func (ic *InventoryControllers) RegisterInventoryRoutes(rg *gin.RouterGroup) {
	inventoryRouter := rg.Group("/")
	inventoryRouter.POST("/register", ic.RegisterInventory)
	inventoryRouter.POST("/:inventoryId/product/add", ic.AddProduct)
	inventoryRouter.GET("/:inventoryId/product/:productId", ic.GetProduct)
}
