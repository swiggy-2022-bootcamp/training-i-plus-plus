package main

import (
	"Inventory-Service/config"
	"Inventory-Service/controller"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/catalog", controller.CreateProduct)
	r.GET("/catalog", controller.GetCatalog)
	r.GET("/catalog/:productId", controller.GetProductById)
	r.PUT("/catalog/:productId", controller.UpdateProductById)
	r.DELETE("/catalog/:productId", controller.DeleteProductbyId)

	portAddress := ":" + strconv.Itoa(config.INVENTORY_SERVICE_SERVER_PORT)
	r.Run(portAddress)

	/*
		inventoryRoutes.GET("/catalog", controller.GetCatalog)
		inventoryRoutes.GET("/catalog/{productId}", controller.GetProductById)
		inventoryRoutes.PUT("/catalog/{productId}", controller.UpdateProductById)
		inventoryRoutes.DELETE("/catalog/{productId}", controller.DeleteProductbyId)
	*/

	// log.Print("Inventory Service: Starting server at port ", config.INVENTORY_SERVICE_SERVER_PORT)
	// http.ListenAndServe(":5002", router)
}
