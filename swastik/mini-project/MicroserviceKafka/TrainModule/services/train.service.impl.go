package services

import (
	"github.com/swastiksahoo153/MicroserviceKafka/TrainModule/models"
	"go.mongodb.org/mongo-driver/mongo"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"errors"
)

type TrainServiceImpl struct {
	traincollection	*mongo.Collection
	ctx 			context.Context
}

func NewTrainService (traincollection *mongo.Collection, ctx context.Context) TrainService{
	return &TrainServiceImpl{
		traincollection:	traincollection,
		ctx:				ctx,
	}
}

func (t *TrainServiceImpl) CreateTrain(train *models.Train) error{
	var seats_available []int
	for i := 1; i<= train.Total_seats; i++{
		seats_available = append(seats_available, i)
	}
	train.Seats_available = seats_available
	_, err := t.traincollection.InsertOne(t.ctx, train)
	return err
}

func (t *TrainServiceImpl) GetTrain(name *string) (*models.Train, error){
	var train *models.Train
	query := bson.D{bson.E{Key:"train_number", Value: name}}
	err := t.traincollection.FindOne(t.ctx, query).Decode(&train)
	return train, err
}

func (t *TrainServiceImpl) GetAll() ([]*models.Train, error){
	var trains []*models.Train
	cursor, err := t.traincollection.Find(t.ctx, bson.D{{}})
	if err != nil{
		return nil, err
	}
	for cursor.Next(t.ctx){
		var train models.Train
		err := cursor.Decode(&train)
		if err != nil {
			return nil, err
		}
		trains = append(trains, &train)
	}
	if err := cursor.Err(); err != nil{
		return nil, err
	}

	cursor.Close(t.ctx)

	if len(trains) == 0 {
		return nil, errors.New("documents not found")
	}
	return trains, nil
}

func (t *TrainServiceImpl) UpdateTrain(train *models.Train) error{
	filter := bson.D{bson.E{Key:"train_number", Value: train.Train_number}}
	update := bson.D{
		bson.E{
			Key:"$set", 
			Value: bson.D{
				bson.E{Key:"train_number", Value: train.Train_number}, 
				bson.E{Key:"train_name", Value: train.Train_name}, 
				bson.E{Key:"source", Value: train.Source}, 
				bson.E{Key:"destination", Value: train.Destination}, 
				bson.E{Key:"total_seats", Value: train.Total_seats},
			}}}

	result,_ := t.traincollection.UpdateOne(t.ctx, filter, update)
	if result.MatchedCount != 1{
		return errors.New("no match found for update")
	}
	return nil
}

func (t *TrainServiceImpl) DeleteTrain(name *string) error{
	// filter := bson.D{bson.E{Key:"user_name", Value: name}}
	// result, _ = t.usercollection.DeleteOne(t.ctx, filter)
	// if result.DeletedCount != 1{
	// 	return errors.New("no match found for delete")
	// } 	
	return nil
}
