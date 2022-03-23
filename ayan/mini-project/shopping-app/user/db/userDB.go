package db

import (
	"context"
	"errors"
	"time"
	"user/domain"

	"go.mongodb.org/mongo-driver/bson"
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
		u.Email(),
		u.Password(),
		u.Name(),
		u.Address(),
		u.Zipcode(),
		u.MobileNo(),
		u.Role(),
	)
	newUser.SetCreatedAt(time.Now())
	newUser.SetUpdatedAt(time.Now())

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

	err := userCollection.FindOne(ctx, bson.M{"email": email}).Decode(dbUser)

	if err != nil {
		return nil, err
	}

	domainUser := domain.NewUser(dbUser.email, dbUser.password, dbUser.name, dbUser.address, dbUser.zipcode, dbUser.mobileNo, dbUser.role)

	return domainUser, errors.New("user does not exist")
}

func (udb userRepositoryDB) UpdateUser(u domain.User) (*domain.User, error) {

	newUser := NewUser(
		u.Email(),
		u.Password(),
		u.Name(),
		u.Address(),
		u.Zipcode(),
		u.MobileNo(),
		u.Role(),
	)
	newUser.SetCreatedAt(time.Now())
	newUser.SetUpdatedAt(time.Now())
	dbUser := User{}

	ctx, cxl := context.WithTimeout(context.Background(), 10*time.Second)
	defer cxl()

	userCollection := Collection(udb.dbClient, "users")

	err := userCollection.FindOneAndReplace(ctx, bson.M{"email": newUser.email}, newUser).Decode(dbUser)

	if err != nil {
		return nil, err
	}

	domainUser := domain.NewUser(dbUser.email, dbUser.password, dbUser.name, dbUser.address, dbUser.zipcode, dbUser.mobileNo, dbUser.role)

	return domainUser, nil
}
