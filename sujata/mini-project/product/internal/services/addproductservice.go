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

type AddProductService interface {
	ValidateRequest(product model.Product) *errors.ServerError
	ProcessRequest(ctx context.Context, product model.Product) *errors.ServerError
}

var addProductServiceStruct AddProductService
var addProductServiceOnce sync.Once

type addProductService struct {
	config *util.RouterConfig
	dao    mongodao.MongoDAO
}

func InitAddProductService(config *util.RouterConfig, dao mongodao.MongoDAO) AddProductService {
	addProductServiceOnce.Do(func() {
		addProductServiceStruct = &addProductService{
			config: config,
			dao:    dao,
		}
	})

	return addProductServiceStruct
}

func GetAddProductService() AddProductService {
	if addProductServiceStruct == nil {
		panic("Add product service not initialised")
	}

	return addProductServiceStruct
}

func (s *addProductService) ValidateRequest(product model.Product) *errors.ServerError {
	if product.Name == "" || product.Description == "" || product.Price == 0 || product.Quantity == 0 {
		log.Error("Either product name, description, price or quantity missing from product")
		return &errors.ParametersMissingError
	}
	return nil
}

func (service *addProductService) ProcessRequest(ctx context.Context, product model.Product) *errors.ServerError {
	err := service.dao.AddProduct(ctx, product)
	if err != nil {
		log.WithField("Error: ", err).Error("an error occurred while inserting product in db")
		return err
	}

	return nil
}
