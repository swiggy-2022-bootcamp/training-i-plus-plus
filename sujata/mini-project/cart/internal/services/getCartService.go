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

type GetCartService interface {
	ProcessRequest(ctx context.Context, email string) (model.Cart, *errors.ServerError)
}

var getCartServiceStruct GetCartService
var getCartServiceOnce sync.Once

type getCartService struct {
	config *util.RouterConfig
	dao    mongodao.MongoDAO
}

func InitGetCartService(config *util.RouterConfig, dao mongodao.MongoDAO) GetCartService {
	getCartServiceOnce.Do(func() {
		getCartServiceStruct = &getCartService{
			config: config,
			dao:    dao,
		}
	})

	return getCartServiceStruct
}

func GetGetCartService() GetCartService {
	if getCartServiceStruct == nil {
		panic("Get cart service not initialised")
	}

	return getCartServiceStruct
}

func (s *getCartService) ProcessRequest(ctx context.Context, email string) (model.Cart, *errors.ServerError) {
	cart, err := s.dao.GetCart(ctx, email)
	if err != nil {
		log.WithField("Error: ", err).Error("an error occurred while getting cart for the user: ", email)
		return cart, err
	}

	return cart, nil
}
