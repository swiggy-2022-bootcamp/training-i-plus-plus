package services

import (
	"context"
	"errors"

	"github.com/go-kafka-microservice/WalletService/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type WalletServicesImpl struct {
	WalletCollection *mongo.Collection
	Ctx              context.Context
}

func NewWalletServiesImpl(wallectCollection *mongo.Collection, ctx context.Context) *WalletServicesImpl {
	return &WalletServicesImpl{
		WalletCollection: wallectCollection,
		Ctx:              ctx,
	}
}

func (ws *WalletServicesImpl) CreateWallet(wallet *models.Wallet) (string, error) {
	wallet.ID = primitive.NewObjectID()
	if _, err := ws.WalletCollection.InsertOne(ws.Ctx, wallet); err != nil {
		return "", err
	}
	return wallet.ID.Hex(), nil
}

func (ws *WalletServicesImpl) AddMoney(walletID primitive.ObjectID, _amount int) error {
	// Mongo Query to filter and update wallet
	filter := bson.D{bson.E{Key: "_id", Value: walletID}}
	inc := bson.D{bson.E{Key: "$inc", Value: bson.D{bson.E{Key: "amount", Value: _amount}}}}

	// Command
	res, err := ws.WalletCollection.UpdateOne(ws.Ctx, filter, inc)
	if err != nil {
		return err
	}
	if res.MatchedCount != 1 {
		return errors.New("No Wallet Found.")
	}
	if res.ModifiedCount != 1 {
		return errors.New("Amount haven't added, please try again.")
	}
	return nil
}

func (ws *WalletServicesImpl) GetStatus(walletID primitive.ObjectID) (*models.Wallet, error) {
	filter := bson.D{bson.E{Key: "_id", Value: walletID}}
	var _wallet models.Wallet
	if err := ws.WalletCollection.FindOne(ws.Ctx, filter).Decode(&_wallet); err != nil {
		return nil, err
	}
	return &_wallet, nil
}

func (ws *WalletServicesImpl) DeductAmount(userId primitive.ObjectID, bill int) error {
	filter := bson.D{bson.E{Key: "user_id", Value: userId}}
	update := bson.D{
		bson.E{Key: "$inc", Value: bson.D{
			bson.E{Key: "amount", Value: -bill},
		}},
	}
	res, err := ws.WalletCollection.UpdateOne(ws.Ctx, filter, update)
	if err != nil {
		return err
	}
	if res.MatchedCount != 1 {
		return errors.New("You have not initiated your wallet.")
	}
	if res.ModifiedCount != 1 {
		return errors.New("Money didn't deduct.")
	}
	return nil
}
