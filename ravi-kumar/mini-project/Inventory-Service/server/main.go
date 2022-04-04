package main

import (
	"Inventory-Service/config"
	"Inventory-Service/controller"
	"Inventory-Service/middleware"
	"Inventory-Service/service"
	"io"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	//logs
	inventorySerivceLogFile, _ := os.Create("inventoryService.log")
	gin.DefaultWriter = io.MultiWriter(os.Stdout, inventorySerivceLogFile)
	service.GetDefaultWriter = &gin.DefaultWriter
	service.InitLoggerService()

	r := gin.Default()

	r.Use(cors.AllowAll())

	r.POST("/catalog", middleware.IfAuthorized(controller.CreateProduct))
	r.GET("/catalog", middleware.IfAuthorized(controller.GetCatalog))
	r.GET("/catalog/:productId", middleware.IfAuthorized(controller.GetProductById))
	r.PUT("/catalog/:productId", middleware.IfAuthorized(controller.UpdateProductById))
	r.DELETE("/catalog/:productId", middleware.IfAuthorized(controller.DeleteProductbyId))
	r.POST("/catalog/:productId/:updateCount", middleware.IfAuthorized(controller.UpdateProductQuantity))
	r.Static("/swaggerui", config.SWAGGER_PATH)

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
