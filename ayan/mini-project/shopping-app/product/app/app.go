package app

import (
	"fmt"

	"product/db"
	"product/docs"
	"product/domain"
	"product/utils/logger"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Start() {

	dbClient := db.NewDbClient()

	productRepo := db.NewProductRepositoryDB(dbClient)
	productService := domain.NewProductService(productRepo)
	productHandlers := ProductHandlers{service: productService}

	productRouter := gin.Default()

	apiRouter := productRouter.Group("/api")

	docs.SwaggerInfo.BasePath = "/api"
	apiRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	productRoutesGroup := apiRouter.Group("/products")

	productRoutesGroup.GET("/", productHandlers.HelloWorldHandler)
	productRoutesGroup.POST("/", productHandlers.Register)
	productRoutesGroup.GET("/:productId", productHandlers.GetProductById)
	productRoutesGroup.PUT("/:productId", productHandlers.UpdateProduct)
	productRoutesGroup.DELETE("/:productId", productHandlers.DeleteProduct)

	productRouter.Run(":8081")
	logger.Info(fmt.Sprintf("Starting server on %s:%s ...", "127.0.0.1", "8081"))
}
