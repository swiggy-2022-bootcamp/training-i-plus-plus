package db

import (
	"context"
	"fmt"
	"time"
	"user/domain"
	"user/utils/errs"
	"user/utils/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepositoryDB struct {
	dbClient *mongo.Client
}

func NewUserRepositoryDB(dbClient *mongo.Client) domain.UserRepositoryDB {
	return &userRepositoryDB{
		dbClient: dbClient,
	}
}

func (udb userRepositoryDB) Save(u domain.User) (*domain.User, *errs.AppError) {

	newUser := NewUser(
		u.Email,
		u.Password,
		u.Name,
		u.Address,
		u.Zipcode,
		u.MobileNo,
		u.Role,
	)
	newUser.SetId(primitive.NewObjectID())
	newUser.SetCreatedAt(time.Now())
	newUser.SetUpdatedAt(time.Now())

	fmt.Println(newUser)

	ctx, cxl := context.WithTimeout(context.Background(), 10*time.Second)
	defer cxl()

	userCollection := Collection(udb.dbClient, "users")
	_, err := userCollection.InsertOne(ctx, newUser)

	if err != nil {
		logger.Error("Error while inserting User into DB : " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from DB")
	}

	return &u, nil
}

func (udb userRepositoryDB) FetchUserByEmail(email string) (*domain.User, *errs.AppError) {

	dbUser := User{}

	ctx, cxl := context.WithTimeout(context.Background(), 10*time.Second)
	defer cxl()

	userCollection := Collection(udb.dbClient, "users")

	err := userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&dbUser)

	if err != nil {
		logger.Error("Error while fetching User from DB : " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from DB")
	}

	domainUser := domain.NewUser(dbUser.Email, dbUser.Password, dbUser.Name, dbUser.Address, dbUser.Zipcode, dbUser.MobileNo, dbUser.Role)

	return domainUser, nil
}

func (udb userRepositoryDB) UpdateUser(u domain.User) (*domain.User, *errs.AppError) {

	ctx, cxl := context.WithTimeout(context.Background(), 10*time.Second)
	defer cxl()

	userCollection := Collection(udb.dbClient, "users")

	currDbUser := User{}
	err := userCollection.FindOne(ctx, bson.M{"email": u.Email}).Decode(&currDbUser)
	if err != nil {
		logger.Error("Error while fetching User from DB : " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from DB")
	}

	newUser := NewUser(
		u.Email,
		u.Password,
		u.Name,
		u.Address,
		u.Zipcode,
		u.MobileNo,
		u.Role,
	)
	newUser.SetCreatedAt(time.Now())
	newUser.SetUpdatedAt(time.Now())
	newUser.SetId(currDbUser.Id)
	dbUser := User{}

	err = userCollection.FindOneAndReplace(ctx, bson.M{"email": newUser.Email}, newUser).Decode(&dbUser)

	if err != nil {
		logger.Error("Error while updating User in DB : " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from DB")
	}

	domainUser := domain.NewUser(dbUser.Email, dbUser.Password, dbUser.Name, dbUser.Address, dbUser.Zipcode, dbUser.MobileNo, dbUser.Role)

	return domainUser, nil
}

func (udb userRepositoryDB) DeleteUserByEmail(email string) *errs.AppError {

	ctx, cxl := context.WithTimeout(context.Background(), 10*time.Second)
	defer cxl()

	userCollection := Collection(udb.dbClient, "users")

	_, err := userCollection.DeleteOne(ctx, bson.M{"email": email})

	if err != nil {
		logger.Error("Error while deleting User from DB : " + err.Error())
		return errs.NewUnexpectedError("Unexpected error from DB")
	}

	return nil
}
