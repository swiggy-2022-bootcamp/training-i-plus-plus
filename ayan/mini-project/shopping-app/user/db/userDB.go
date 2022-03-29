package db

import (
	"context"
	"fmt"
	"time"
	"user/domain"

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

func (udb userRepositoryDB) Save(u domain.User) (*domain.User, error) {

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
		return nil, err
	}

	return &u, nil
}

func (udb userRepositoryDB) FindUserByEmail(email string) (*domain.User, error) {

	dbUser := User{}

	ctx, cxl := context.WithTimeout(context.Background(), 10*time.Second)
	defer cxl()

	userCollection := Collection(udb.dbClient, "users")

	err := userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&dbUser)

	fmt.Println("(udb userRepositoryDB) FindUserByEmail : ", dbUser, err)

	if err != nil {
		return nil, err
	}

	domainUser := domain.NewUser(dbUser.Email, dbUser.Password, dbUser.Name, dbUser.Address, dbUser.Zipcode, dbUser.MobileNo, dbUser.Role)

	return domainUser, nil
}

func (udb userRepositoryDB) UpdateUser(u domain.User) (*domain.User, error) {

	ctx, cxl := context.WithTimeout(context.Background(), 10*time.Second)
	defer cxl()

	userCollection := Collection(udb.dbClient, "users")

	currDbUser := User{}
	err := userCollection.FindOne(ctx, bson.M{"email": u.Email}).Decode(&currDbUser)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	domainUser := domain.NewUser(dbUser.Email, dbUser.Password, dbUser.Name, dbUser.Address, dbUser.Zipcode, dbUser.MobileNo, dbUser.Role)

	return domainUser, nil
}

func (udb userRepositoryDB) DeleteUserByEmail(email string) error {

	ctx, cxl := context.WithTimeout(context.Background(), 10*time.Second)
	defer cxl()

	userCollection := Collection(udb.dbClient, "users")

	_, err := userCollection.DeleteOne(ctx, bson.M{"email": email})

	fmt.Println("(udb userRepositoryDB) DeleteUserByEmail : ", err)

	if err != nil {
		return err
	}

	return nil
}
