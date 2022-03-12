package controller

import (
	"github.com/gin-gonic/gin"
	services "github.com/sachinsom93/shopping-cart/services/inventory"
)

type InventoryController struct {
	InventoryService *services.InventoryServiceImpl
}

// Function to create new instance of inventory controller
func NewInventoryController(inventoryService *services.InventoryServiceImpl) *InventoryController {
	return &InventoryController{
		InventoryService: inventoryService,
	}
}

// Function to register a new inventory
func (ic *InventoryController) RegisterInventory(gctx *gin.Context) {
	gctx.JSON(200, "")
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
