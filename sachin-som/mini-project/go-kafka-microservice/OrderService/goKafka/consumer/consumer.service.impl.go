package goKafka

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/go-kafka-microservice/OrderService/models"
	pb "github.com/go-kafka-microservice/WalletProto"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type GoKafkaServicesImpl struct {
	Consumer          *kafka.Reader
	OrderCollection   *mongo.Collection
	WalletProtoClient pb.WalletServiceClient
	Ctx               context.Context
}

func NewGokafkaServiceImpl(consumer *kafka.Reader, orderCollection *mongo.Collection, walletProtoClient pb.WalletServiceClient, ctx context.Context) *GoKafkaServicesImpl {
	return &GoKafkaServicesImpl{
		Consumer:          consumer,
		OrderCollection:   orderCollection,
		WalletProtoClient: walletProtoClient,
		Ctx:               ctx,
	}
}
func (ks *GoKafkaServicesImpl) StoreOrders(topic string) error {

	// 1. Get Product from ordered_products topic
	// 2. Create Instance of Order Model
	// 3. Store Order to database
	// 4. Deduct Wallet Amount (TODO)
	// 5. Notify Inventory Owner (TODO)
	// 5. Create Bill (TODO)
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := ks.Consumer.ReadMessage(ks.Ctx)
		if err != nil {
			return err
		}
		var _userProduct models.UserProduct
		p := []byte(msg.Value)
		err = json.Unmarshal(p, &_userProduct)
		if err != nil {
			return err
		}

		// Create New Order Instance
		_product := models.Product{
			ID:          _userProduct.ID,
			ProductName: _userProduct.ProductName,
			Description: _userProduct.Description,
			Quantity:    "1",
			Price:       _userProduct.Price,
			Ratings:     _userProduct.Ratings,
			ImageUrl:    _userProduct.ImageUrl,
		}
		_order := models.Order{
			OrderID:       primitive.NewObjectID(),
			OrderCart:     _product,
			OrderedAt:     time.Now(),
			Bill:          _product.Price,
			Discount:      "",
			PaymentMethod: "Wallet",
			Status:        "initiated",
			UserID:        _userProduct.UserID,
		}
		if _, err := ks.OrderCollection.InsertOne(ks.Ctx, _order); err != nil {
			return err
		}

		// Check available wallet amount
		userInfo := &pb.UserInfo{
			UserId: _userProduct.UserID.Hex(),
		}
		res, err := ks.WalletProtoClient.CheckAmount(ks.Ctx, userInfo)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		amount := res.Amount
		bill, _ := strconv.Atoi(_order.Bill)
		if amount < int64(bill) {
			filter := bson.D{bson.E{Key: "_id", Value: _order.OrderID}}
			update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "status", Value: "failed"}}}}
			ks.OrderCollection.UpdateOne(ks.Ctx, filter, update)
			return errors.New("Not Sufficient Amount.")
			// TODO: Need to notify client (buyer) through Email service for order failure.
		}

		// Deduct wallet amount
		deductReq := &pb.DeductRequest{
			UserId: _userProduct.UserID.Hex(),
			Bill:   int64(bill),
		}
		_, err = ks.WalletProtoClient.DeductAmount(ks.Ctx, deductReq)
		if err != nil {
			fmt.Println(err.Error())
		}
		// Change order status
		filter := bson.D{bson.E{Key: "_id", Value: _order.OrderID}}
		update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "status", Value: "paid"}}}}
		ks.OrderCollection.UpdateOne(ks.Ctx, filter, update)
		// TODO: Need to notify client (buyer) for successfull order
	}
}
