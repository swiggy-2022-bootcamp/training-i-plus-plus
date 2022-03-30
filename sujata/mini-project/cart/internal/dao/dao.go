package mongodao

import (
	model "cart/internal/dao/models"
	"cart/internal/errors"
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	log "github.com/sirupsen/logrus"
)

type MongoDAO interface {
	AddProduct(ctx context.Context, cartProduct model.CartProduct, email string) *errors.ServerError
	DeleteProduct(ctx context.Context, productId string, email string) *errors.ServerError
	UpdateProductQuantity(ctx context.Context) *errors.ServerError
	GetCart(ctx context.Context, email string) (model.Cart, *errors.ServerError)
	DeleteCart(ctx context.Context, email string) *errors.ServerError
}

type mongoDAO struct {
	client *mongo.Client
}

var mongoDao MongoDAO
var mongoDaoOnce sync.Once

func InitMongoDAO(client *mongo.Client) MongoDAO {
	mongoDaoOnce.Do(func() {
		mongoDao = &mongoDAO{
			client: client,
		}
	})

	return mongoDao
}

func GetMongoDAO() MongoDAO {
	if mongoDao == nil {
		panic("Mongo DAO not initialised")
	}

	return mongoDao
}

// AddProduct add product details to the cart of the user
func (dao *mongoDAO) AddProduct(ctx context.Context, cartProduct model.CartProduct, email string) *errors.ServerError {
	cartCollection := dao.client.Database("shopKart").Collection("cart")

	filter := bson.M{"email": email}

	// create an instance of an options and set the desired options
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	update := bson.M{
		"$push": bson.M{"products": cartProduct},
	}

	result := cartCollection.FindOneAndUpdate(ctx, filter, update, &opt)
	if result.Err() != nil {
		log.WithError(result.Err()).Error("an error occurred while adding the product to the cart")
		return &errors.InternalError
	}

	return nil
}

// DeleteProduct removes the product details from the cart of the user
func (dao *mongoDAO) DeleteProduct(ctx context.Context, productId string, email string) *errors.ServerError {
	cartCollection := dao.client.Database("shopKart").Collection("cart")

	filter := bson.M{"email": email}
	update := bson.M{
		"$pull": bson.M{"products": bson.M{"productId": productId}},
	}
	upsert := false
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}
	result, err := cartCollection.UpdateOne(ctx, filter, update, &opt)
	if err != nil {
		log.WithField("Error: ", err).Error("an error occurred while deleting the product")
		return &errors.InternalError
	}

	if result.MatchedCount == 0 {
		log.Info("No product with productId: ", productId, " in the cart was found for user: ", email)
		return &errors.InternalError
	}

	if result.ModifiedCount == 0 {
		log.Error("No product product deleted for user: ", email)
		return &errors.InternalError
	}

	return nil
}

// UpdateProductQuantity update the product quantity in the cart of the user
func (dao *mongoDAO) UpdateProductQuantity(ctx context.Context) *errors.ServerError {
	return nil
}

// GetCart returns products in cart of the user along with total price
func (dao *mongoDAO) GetCart(ctx context.Context, email string) (model.Cart, *errors.ServerError) {
	cartCollection := dao.client.Database("shopKart").Collection("cart")

	filter := bson.M{"email": email}
	result := cartCollection.FindOne(ctx, filter)

	cart := model.Cart{}
	err := result.Decode(&cart)
	if err != nil {
		log.WithError(err).Error("an error occurred while decoding the cart details for user: ", email)
		return cart, &errors.InternalError
	}

	return cart, nil
}

func (dao *mongoDAO) DeleteCart(ctx context.Context, email string) *errors.ServerError {
	cartCollection := dao.client.Database("shopKart").Collection("cart")

	filter := bson.M{"email": email}
	result, err := cartCollection.DeleteOne(ctx, filter)
	if err != nil {
		log.WithError(err).Error("an error occurred while deleting cart for user: ", email)
		return &errors.InternalError
	}

	if result.DeletedCount == 0 {
		log.Error("Deleted documents: 0, expected 1")
		return &errors.InternalError
	}

	log.Info("Cart deleted successfully for user: ", email)
	return nil
}
