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
	var product models.Product
	if err := gctx.ShouldBindJSON(&product); err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := ic.InventoryService.AddProduct(&product); err != nil {
		gctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	gctx.JSON(http.StatusCreated, gin.H{"message": "Product added to inventory."})
}

// Function to remove a specific product from inventory
func (ic *InventoryController) RemoveProduct(gctx *gin.Context) {
	productID := gctx.Param("productId")
	inventoryID := gctx.Param("inventoryId")
	var err error
	if inventoryID == "" {
		gctx.JSON(http.StatusBadRequest, gin.H{"message": "Please provide inventory id."})
		return
	}
	if productID == "" {
		gctx.JSON(http.StatusBadRequest, gin.H{"message": "Please provide a product id."})
		return
	}
	if err = ic.InventoryService.RemoveProduct(inventoryID, productID); err != nil {
		gctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	gctx.JSON(http.StatusOK, gin.H{"message": "Product removed from inventory."})
}

// Function to update product by given product instance
func (ic *InventoryController) UpdateProduct(gctx *gin.Context) {
	gctx.JSON(200, "")
}

// Function to get a specific product item
func (ic *InventoryController) GetProduct(gctx *gin.Context) {
	productID := gctx.Param("productId")
	inventoryID := gctx.Param("inventoryId")
	var err error
	if inventoryID == "" {
		gctx.JSON(http.StatusBadRequest, gin.H{"message": "Please provide inventory id."})
		return
	}
	if productID == "" {
		gctx.JSON(http.StatusBadRequest, gin.H{"message": "Please provide a product id."})
		return
	}
	var product *models.Product
	if product, err = ic.InventoryService.GetProduct(inventoryID, productID); err != nil {
		gctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	gctx.JSON(http.StatusOK, product)
}

// Function to get all product of an inventory
func (ic *InventoryController) GetAllProducts(gctx *gin.Context) {
	var inventoryId string
	var products []models.Product
	var err error
	inventoryId = gctx.Param("inventoryId")
	if products, err = ic.InventoryService.GetAllProducts(inventoryId); err != nil {
		gctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	gctx.JSON(http.StatusOK, products)
}

func (ic *InventoryController) RegisterInventoryRoutes(rg *gin.RouterGroup) {
	inventoryRoute := rg.Group("/inventory")
	inventoryRoute.POST("/register", ic.RegisterInventory)
	inventoryRoute.POST("/addproduct", ic.AddProduct)
	inventoryRoute.GET("/get/:inventoryId/:productId", ic.GetProduct)
	inventoryRoute.DELETE("/remove/:inventoryId/:productId", ic.GetProduct)
	inventoryRoute.GET("/getall/:inventoryId", ic.GetAllProducts)
	inventoryRoute.PATCH("/update", ic.UpdateProduct)
}
