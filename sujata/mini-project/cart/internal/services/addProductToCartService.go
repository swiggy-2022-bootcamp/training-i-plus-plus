package services

import (
	mongodao "cart/internal/dao"
	model "cart/internal/dao/models"
	"cart/internal/errors"
	"cart/util"
	"context"
	"sync"

	log "github.com/sirupsen/logrus"
)

type AddProductToCartService interface {
	ValidateRequest(product model.CartProduct) *errors.ServerError
	ProcessRequest(ctx context.Context, email string, product model.CartProduct) *errors.ServerError
}

var addProductToCartServiceStruct AddProductToCartService
var addProductToCartServiceOnce sync.Once

type addProductToCartService struct {
	config *util.RouterConfig
	dao    mongodao.MongoDAO
}

func InitAddProductToCartService(config *util.RouterConfig, dao mongodao.MongoDAO) AddProductToCartService {
	addProductToCartServiceOnce.Do(func() {
		addProductToCartServiceStruct = &addProductToCartService{
			config: config,
			dao:    dao,
		}
	})

	return addProductToCartServiceStruct
}

func GetAddProductToCartService() AddProductToCartService {
	if addProductToCartServiceStruct == nil {
		panic("Add product to cart service not initialised")
	}

	return addProductToCartServiceStruct
}

func (s *addProductToCartService) ValidateRequest(product model.CartProduct) *errors.ServerError {
	if product.ProductId == "" || product.Quantity == 0 {
		log.Error("Either product id missing or product quantity missing from the product")
		return &errors.ParametersMissingError
	}

	return nil
}

func (s *addProductToCartService) ProcessRequest(ctx context.Context, email string, product model.CartProduct) *errors.ServerError {
	err := s.dao.AddProduct(ctx, product, email)
	if err != nil {
		log.WithField("Error: ", err).Error("an error occurred while inserting product in the cart for user: ", email)
		return err
	}

	return nil
}
