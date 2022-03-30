package main

import (
	"Inventory-Service/config"
	"Inventory-Service/controller"
	"Inventory-Service/middleware"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/catalog", middleware.IfAuthorized(controller.CreateProduct))
	r.GET("/catalog", middleware.IfAuthorized(controller.GetCatalog))
	r.GET("/catalog/:productId", middleware.IfAuthorized(controller.GetProductById))
	r.PUT("/catalog/:productId", middleware.IfAuthorized(controller.UpdateProductById))
	r.DELETE("/catalog/:productId", middleware.IfAuthorized(controller.DeleteProductbyId))
	r.POST("/catalog/:productId/:updateCount", middleware.IfAuthorized(controller.UpdateProductQuantity))

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
