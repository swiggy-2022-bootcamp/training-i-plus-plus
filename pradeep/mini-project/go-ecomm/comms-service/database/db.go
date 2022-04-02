package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type DbHandler struct {
	Collection *mongo.Collection
	Ctx        context.Context
}

func NewDbHandler(collection *mongo.Collection, ctx context.Context) *DbHandler {
	return &DbHandler{
		Collection: collection,
		Ctx:        ctx,
	}
}
