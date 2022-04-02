package services

import (
	"context"
	mongodao "order/internal/dao"
	model "order/internal/dao/models"
	"order/internal/errors"
	"order/util"
	"sync"
)

type GetOrderService interface {
	ProcessRequest(ctx context.Context, email string) (model.AllOrders, *errors.ServerError)
}

var getOrderServiceStruct GetOrderService
var getOrderServiceOnce sync.Once

type getOrderService struct {
	config *util.RouterConfig
	dao    mongodao.MongoDAO
}

func InitGetOrderService(config *util.RouterConfig, dao mongodao.MongoDAO) GetOrderService {
	getOrderServiceOnce.Do(func() {
		getOrderServiceStruct = &getOrderService{
			config: config,
			dao:    dao,
		}
	})

	return getOrderServiceStruct
}

func GetGetOrderService() GetOrderService {
	if getOrderServiceStruct == nil {
		panic("Get order service not initialised")
	}

	return getOrderServiceStruct
}

func (s *getOrderService) ProcessRequest(ctx context.Context, email string) (model.AllOrders, *errors.ServerError) {
	return s.dao.GetOrders(ctx, email)
}
