package services

import (
	mongodao "cart/internal/dao"
	"cart/internal/errors"
	"cart/util"
	"context"
	"sync"

	log "github.com/sirupsen/logrus"
)

type DeleteProductFromCartService interface {
	ValidateRequest(productId string) *errors.ServerError
	ProcessRequest(ctx context.Context, productId string, email string) *errors.ServerError
}

var deleteProductFromCartServiceStruct DeleteProductFromCartService
var deleteProductFromCartServiceOnce sync.Once

type deleteProductFromCartService struct {
	config *util.RouterConfig
}

func InitDeleteProductFromCartService(config *util.RouterConfig) DeleteProductFromCartService {
	deleteProductFromCartServiceOnce.Do(func() {
		deleteProductFromCartServiceStruct = &deleteProductFromCartService{
			config: config,
		}
	})

	return deleteProductFromCartServiceStruct
}

func GetDeleteProductFromCartService() DeleteProductFromCartService {
	if deleteProductFromCartServiceStruct == nil {
		panic("Delete product from cart service not initialised")
	}

	return deleteProductFromCartServiceStruct
}

func (s *deleteProductFromCartService) ValidateRequest(productId string) *errors.ServerError {
	if productId == "" {
		log.Error("productId send in the request as query parameter is empty")
		return &errors.QueryParamMissingError
	}

	if productId[0] == '"' {
		log.Error("malformed query param 'productId' in the request")
		return &errors.MalformedQueryParamError
	}

	return nil
}

func (s *deleteProductFromCartService) ProcessRequest(ctx context.Context, productId string, email string) *errors.ServerError {
	dao := mongodao.GetMongoDAO()

	err := dao.DeleteProduct(ctx, productId, email)
	if err != nil {
		log.WithField("Error: ", err).Error("an error occurred while deleting the product: ", productId, " from the cart for user: ", email)
		return err
	}

	return nil
}
