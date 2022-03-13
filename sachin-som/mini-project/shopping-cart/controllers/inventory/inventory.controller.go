package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sachinsom93/shopping-cart/models"
	services "github.com/sachinsom93/shopping-cart/services/inventory"
)

type InventoryController struct {
	InventoryService services.InventoryServices
}

// Function to create new instance of inventory controller
func NewInventoryController(inventoryService services.InventoryServices) *InventoryController {
	return &InventoryController{
		InventoryService: inventoryService,
	}
}

// Function to register a new inventory
func (ic *InventoryController) RegisterInventory(gctx *gin.Context) {
	var inventory models.Inventory
	if err := gctx.ShouldBindJSON(&inventory); err != nil {
		gctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	if err := ic.InventoryService.RegisterInventory(&inventory); err != nil {
		gctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	gctx.JSON(http.StatusCreated, gin.H{"message": "Inventory registered succesfully."})
}

// Function to add a specific product into inventory
func (ic *InventoryController) AddProduct(gctx *gin.Context) {
	gctx.JSON(200, "")
}

// Function to remove a specific product from inventory
func (ic *InventoryController) RemoveProduct(gctx *gin.Context) {
	gctx.JSON(200, "")
}

// Function to update product by given product instance
func (ic *InventoryController) UpdateProduct(gctx *gin.Context) {
	gctx.JSON(200, "")
}

// Function to get a specific product item
func (ic *InventoryController) GetProduct(gctx *gin.Context) {
	gctx.JSON(200, "")
}

// Function to get all product of an inventory
func (ic *InventoryController) GetAllProducts(gctx *gin.Context) {
	gctx.JSON(200, "")
}

func (ic *InventoryController) RegisterInventoryRoutes(rg *gin.RouterGroup) {
	inventoryRoute := rg.Group("/inventory")
	inventoryRoute.POST("/register", ic.RegisterInventory)
	inventoryRoute.GET("/get/:productId", ic.GetProduct)
	inventoryRoute.GET("/getall", ic.GetAllProducts)
	inventoryRoute.PATCH("/update", ic.UpdateProduct)
}
