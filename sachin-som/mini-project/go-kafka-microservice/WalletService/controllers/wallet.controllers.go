package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	pb "github.com/go-kafka-microservice/WalletProto"
	"github.com/go-kafka-microservice/WalletService/models"
	"github.com/go-kafka-microservice/WalletService/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WalletControllers struct {
	pb.UnimplementedWalletServiceServer
	WalletServices services.WalletServices
}

func NewWalletControllers(ws services.WalletServices) *WalletControllers {
	return &WalletControllers{
		WalletServices: ws,
	}
}

func (wc *WalletControllers) CreateWallet(g *gin.Context) {
	var _wallet models.Wallet
	if err := g.ShouldBindJSON(&_wallet); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	_walletID, err := wc.WalletServices.CreateWallet(&_wallet)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	g.JSON(http.StatusCreated, gin.H{"wallet_id": _walletID})
}

func (wc *WalletControllers) AddMoney(g *gin.Context) {
	_walletObjId, _ := primitive.ObjectIDFromHex(g.Param("walletId"))
	if _walletObjId == primitive.NilObjectID {
		g.JSON(http.StatusBadRequest, gin.H{"message": "Please provide wallet Id."})
		return
	}
	var _wallet models.Wallet
	if err := g.ShouldBindJSON(&_wallet); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := wc.WalletServices.AddMoney(_walletObjId, _wallet.Amount); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	g.JSON(http.StatusOK, gin.H{"message": "Wallet Updated Successfully."})
}

func (wc *WalletControllers) GetStatus(g *gin.Context) {
	_walletObjId, _ := primitive.ObjectIDFromHex(g.Param("walletId"))
	if _walletObjId == primitive.NilObjectID {
		g.JSON(http.StatusBadRequest, gin.H{"message": "Please provide wallet Id."})
		return
	}
	_wallet, err := wc.WalletServices.GetStatus(_walletObjId)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	g.JSON(http.StatusOK, gin.H{"message": _wallet})
}

func (wc *WalletControllers) RegisterWalletRoutes(rg *gin.RouterGroup) {
	walletRouter := rg.Group("/")
	walletRouter.POST("/create", wc.CreateWallet)
	walletRouter.PATCH("/:walletId/add", wc.AddMoney)
	walletRouter.GET("/:walletId", wc.GetStatus)
}

/*
	Wallet Proto Service implementationis.
*/
func (wc *WalletControllers) CheckAmount(ctx context.Context, in *pb.UserInfo) (*pb.AmountResponse, error) {
	userObjId, err := primitive.ObjectIDFromHex(in.UserId)
	if err != nil {
		return nil, err
	}
	wallet, err := wc.WalletServices.GetStatus(userObjId)
	if err != nil {
		return nil, err
	}
	return &pb.AmountResponse{
		UserId: wallet.UserID.Hex(),
		Amount: int64(wallet.Amount),
	}, nil
}

func (wc *WalletControllers) DeductAmount(ctx context.Context, in *pb.DeductRequest) (*pb.ResponseMessage, error) {
	userObjId, err := primitive.ObjectIDFromHex(in.UserId)
	if err != nil {
		return nil, err
	}
	if err := wc.WalletServices.DeductAmount(userObjId, int(in.Bill)); err != nil {
		return nil, err
	}
	return &pb.ResponseMessage{
		Message: "Amount Deducted From wallet succesfully.",
		Success: true,
	}, err
}
