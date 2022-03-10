package main

import (
	"mini-project/bff/config"
	"mini-project/bff/controller"
	"mini-project/bff/service"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	stationSvc        service.StationService       = service.New()
	stationController controller.StationController = controller.New(stationSvc)
)

func main() {
	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	config.SetupLogPath()

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(config.LoggingMiddleware())
	// TODO hide authorization headers
	router.Use(gindump.Dump())
	// router.Use(config.BasicAuthMiddleware())

	authorized := router.Group("/", config.BasicAuthMiddleware())

	// No authentication
	router.GET("/stations", func(ctx *gin.Context) {
		ctx.JSON(200, stationController.RetrieveAllStations())
	})

	// Has authentication
	authorized.POST("/station", func(ctx *gin.Context) {
		ctx.JSON(200, stationController.AddStation(ctx))
	})

	router.Run()
}
