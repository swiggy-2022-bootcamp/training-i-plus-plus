package app

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swiggy-2022-bootcamp/training-i-plus-plus/murtaza/mini-project/inventory/docs"
	"github.com/swiggy-2022-bootcamp/training-i-plus-plus/murtaza/mini-project/inventory/domain"
	"github.com/swiggy-2022-bootcamp/training-i-plus-plus/murtaza/mini-project/inventory/infra"
	"github.com/swiggy-2022-bootcamp/training-i-plus-plus/murtaza/mini-project/inventory/utils/logger"
)

type Routes struct {
	router *gin.Engine
}

func Start() {

	itemMongoRepository := infra.NewItemMongoRepository()

	itemHandler := ItemHandler{
		itemService: domain.NewItemService(itemMongoRepository),
	}

	r := Routes{
		router: gin.Default(),
	}

	docs.SwaggerInfo.BasePath = "/api/v1"
	r.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.router.Group("/api")
	v1 := api.Group("/v1")

	items := v1.Group("/items")

	//items.GET("/", itemHandler.getAllUsers)
	items.GET("/:itemId", itemHandler.getItemByItemId)
	items.GET("/", itemHandler.getItemByItemName)
	items.DELETE("/:itemId", itemHandler.deleteItem)
	items.PUT("/:itemId", itemHandler.updateItem)
	items.PATCH("/:itemId", itemHandler.updateQuantity)
	items.POST("/", itemHandler.createItem)
	err := r.router.Run(":8090")
	if err != nil {
		logger.Fatal("Unable to start item service")
	}
}
