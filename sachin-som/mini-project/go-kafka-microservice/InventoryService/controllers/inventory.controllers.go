package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kafka-microservice/InventoryService/services"
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
	gctx.JSON(200, nil)
}

func (ic *InventoryControllers) AddProduct(gctx *gin.Context) {
	gctx.JSON(200, nil)
}

func (ic *InventoryControllers) GetProduct(gctx *gin.Context) {
	gctx.JSON(200, nil)
}

func (ic *InventoryControllers) RegisterInventoryRoutes(rg *gin.RouterGroup) {
	inventoryRouter := rg.Group("/inventory")
	inventoryRouter.POST("/regiter", ic.RegisterInventory)
	inventoryRouter.POST("/product/add", ic.AddProduct)
}
