package db

import (
	"context"
	"product/domain"
	"product/utils/errs"
	"product/utils/logger"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type productRepositoryDB struct {
	dbClient *mongo.Client
}

func NewProductRepositoryDB(dbClient *mongo.Client) domain.ProductRepositoryDB {
	return &productRepositoryDB{
		dbClient: dbClient,
	}
}

func (pdb productRepositoryDB) Save(u domain.Product) (*domain.Product, *errs.AppError) {

	newProduct := NewProduct(
		u.Name,
		u.Description,
		u.Quantity,
		u.Price,
	)
	newProduct.Id = primitive.NewObjectID().Hex()

	ctx, cxl := context.WithTimeout(context.Background(), 10*time.Second)
	defer cxl()

	productCollection := Collection(pdb.dbClient, "products")
	_, err := productCollection.InsertOne(ctx, newProduct)

	if err != nil {
		logger.Error("Error while inserting Product into DB : " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from DB")
	}
	u.Id = newProduct.Id

	return &u, nil
}

func (pdb productRepositoryDB) FetchProductById(id string) (*domain.Product, *errs.AppError) {

	dbProduct := Product{}

	ctx, cxl := context.WithTimeout(context.Background(), 10*time.Second)
	defer cxl()

	productCollection := Collection(pdb.dbClient, "products")

	err := productCollection.FindOne(ctx, bson.M{"id": id}).Decode(&dbProduct)

	if err != nil {
		logger.Error("Error while fetching Product from DB : " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from DB")
	}

	domainProduct := domain.NewProduct(dbProduct.Name, dbProduct.Description, dbProduct.Quantity, dbProduct.Price)
	domainProduct.Id = dbProduct.Id

	return domainProduct, nil
}

func (pdb productRepositoryDB) UpdateProduct(id string, u domain.Product) (*domain.Product, *errs.AppError) {

	ctx, cxl := context.WithTimeout(context.Background(), 10*time.Second)
	defer cxl()

	productCollection := Collection(pdb.dbClient, "products")

	currDbProduct := Product{}
	err := productCollection.FindOne(ctx, bson.M{"id": id}).Decode(&currDbProduct)
	if err != nil {
		logger.Error("Error while fetching Product from DB : " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from DB")
	}

	newProduct := NewProduct(
		u.Name,
		u.Description,
		u.Quantity,
		u.Price,
	)
	newProduct.CreatedAt = time.Now()
	newProduct.UpdatedAt = time.Now()
	newProduct.mongoId = currDbProduct.mongoId
	newProduct.Id = currDbProduct.Id
	dbProduct := Product{}

	err = productCollection.FindOneAndReplace(ctx, bson.M{"id": id}, newProduct).Decode(&dbProduct)

	if err != nil {
		logger.Error("Error while updating Product in DB : " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from DB")
	}

	domainProduct := domain.NewProduct(dbProduct.Name, dbProduct.Description, dbProduct.Quantity, dbProduct.Price)
	domainProduct.Id = dbProduct.Id

	return domainProduct, nil
}

func (pdb productRepositoryDB) DeleteProductById(id string) *errs.AppError {

	ctx, cxl := context.WithTimeout(context.Background(), 10*time.Second)
	defer cxl()

	productCollection := Collection(pdb.dbClient, "products")

	_, err := productCollection.DeleteOne(ctx, bson.M{"id": id})

	if err != nil {
		logger.Error("Error while deleting Product from DB : " + err.Error())
		return errs.NewUnexpectedError("Unexpected error from DB")
	}

	return nil
}
