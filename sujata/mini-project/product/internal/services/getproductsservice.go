package services

import (
	"context"
	"product/internal/dao/mongodao"
	model "product/internal/dao/mongodao/models"
	"product/internal/errors"
	"product/util"
	"sync"

	log "github.com/sirupsen/logrus"
)

type GetProductsService interface {
	ProcessRequest(ctx context.Context, productIds model.ProductIds) ([]model.Product, *errors.ServerError)
}

var getProductsServiceStruct GetProductsService
var getProductsServiceOnce sync.Once

type getProductsService struct {
	config *util.RouterConfig
	dao    mongodao.MongoDAO
}

func InitGetProductsService(config *util.RouterConfig, dao mongodao.MongoDAO) GetProductsService {
	getProductsServiceOnce.Do(func() {
		getProductsServiceStruct = &getProductsService{
			config: config,
			dao:    dao,
		}
	})

	return getProductsServiceStruct
}

func GetGetProductsService() GetProductsService {
	if getProductsServiceStruct == nil {
		panic("Get product service not initialised")
	}

	return getProductsServiceStruct
}

func (service *getProductsService) ProcessRequest(ctx context.Context, productIds model.ProductIds) ([]model.Product, *errors.ServerError) {
	var products []model.Product

	for idx := 0; idx < len(productIds.Ids); idx++ {
		id := productIds.Ids[idx]
		product, err := service.dao.GetProduct(ctx, id)
		if err != nil {
			log.WithField("Error: ", err).Error("an error occurred while getting the product from db")
			return products, err
		}

		products = append(products, product)
	}

	return products, nil
}
