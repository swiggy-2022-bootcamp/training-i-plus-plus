package repository

import (
	"context"
	"sanitaria-microservices/patientModule/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type PatientRepository struct {
	MongoCollection *mongo.Collection
}

type MongoOperations interface {
	InsertPatient(ctx context.Context, patient models.Patient)
}

func (p *PatientRepository) InsertPatient(ctx context.Context, patient models.Patient) (*mongo.InsertOneResult, error) {
	result, err := p.MongoCollection.InsertOne(ctx, patient)
	if err != nil {
		return nil, err
	}
	return result, nil
}
