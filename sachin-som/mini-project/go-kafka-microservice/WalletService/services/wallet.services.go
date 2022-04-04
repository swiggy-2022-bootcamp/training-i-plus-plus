package services

import (
	"github.com/go-kafka-microservice/WalletService/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WalletServices interface {
	CreateWallet(*models.Wallet) (string, error)
	AddMoney(primitive.ObjectID, int) error
	GetStatus(primitive.ObjectID) (*models.Wallet, error)
	GetStatusByUserId(primitive.ObjectID) (*models.Wallet, error)
	DeductAmount(primitive.ObjectID, int) error
}
