package services

import (
	"context"
	mongodao "order/internal/dao"
	model "order/internal/dao/models"
	"order/internal/errors"
	"order/util"
	"sync"

	log "github.com/sirupsen/logrus"
)

type SetOrderStatusService interface {
	ValidateRequest(ctx context.Context, orderInfo model.OrderInfo) *errors.ServerError
	ProcessRequest(ctx context.Context, orderInfo model.OrderInfo) *errors.ServerError
}

var setOrderStatusServiceStruct SetOrderStatusService
var setOrderStatusServiceOnce sync.Once

type setOrderStatusService struct {
	config *util.RouterConfig
	dao    mongodao.MongoDAO
}

func InitSetOrderStatusService(config *util.RouterConfig, dao mongodao.MongoDAO) SetOrderStatusService {
	setOrderStatusServiceOnce.Do(func() {
		setOrderStatusServiceStruct = &setOrderStatusService{
			config: config,
			dao:    dao,
		}
	})

	return setOrderStatusServiceStruct
}

func GetSetOrderStatusService() SetOrderStatusService {
	if setOrderStatusServiceStruct == nil {
		panic("Set order status service not initialised")
	}

	return setOrderStatusServiceStruct
}

func (s *setOrderStatusService) ValidateRequest(ctx context.Context, orderInfo model.OrderInfo) *errors.ServerError {
	orderStatus := orderInfo.OrderStatus

	if !(orderStatus == model.CANCELLED || orderStatus == model.DISPATCHED || orderStatus == model.ORDER_COMPLETED || orderStatus == model.OUT_FOR_DELIVERY) {
		log.Error("invalid order status in the request, can not set status to: ", orderStatus)
		return &errors.BadRequest
	}

	return nil
}

func (s *setOrderStatusService) ProcessRequest(ctx context.Context, orderInfo model.OrderInfo) *errors.ServerError {

	return s.dao.SetOrderStatus(ctx, orderInfo)
}
