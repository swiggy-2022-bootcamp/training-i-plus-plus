package goKafka

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-kafka-microservice/ListingService/models"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
)

type GoKafkaServicesImpl struct {
	Consumer          *kafka.Reader
	ProductCollection *mongo.Collection
	Ctx               context.Context
}

func NewGokafkaServiceImpl(consumer *kafka.Reader, productCollection *mongo.Collection, ctx context.Context) *GoKafkaServicesImpl {
	return &GoKafkaServicesImpl{
		Consumer:          consumer,
		ProductCollection: productCollection,
		Ctx:               ctx,
	}
}
func (ks *GoKafkaServicesImpl) StoreOrders(topic string) error {

	// 1. Get Product from ordered_products topic
	// 2. Create Instance of Order Model
	// 3. Store Order to database
	// 4. Notify Inventory Owner
	// 5. Create Bill
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := ks.Consumer.ReadMessage(ks.Ctx)
		if err != nil {
			return err
		}
		var _product models.Product
		p := []byte(msg.Value)
		err = json.Unmarshal(p, &_product)
		fmt.Println(_product)
		if err != nil {
			return err
		}
		if _, err := ks.ProductCollection.InsertOne(ks.Ctx, _product); err != nil {
			return err
		}
	}
}
