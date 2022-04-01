package services

import (
	"github.com/swastiksahoo153/MicroserviceKafka/TicketModule/models"
	kf "github.com/swastiksahoo153/MicroserviceKafka/TicketModule/kafkaProducer"
	"go.mongodb.org/mongo-driver/mongo"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"errors"
	"fmt"
	"time"
)

type TicketServiceImpl struct {
	ticketcollection	*mongo.Collection
	ctx 				context.Context
}

func NewTicketService (ticketcollection *mongo.Collection, ctx context.Context) TicketService{
	return &TicketServiceImpl{
		ticketcollection: 	ticketcollection,
		ctx:				ctx,
	}
}

func (t *TicketServiceImpl) CreateTicket(ticket *models.Ticket) error{
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, err := t.ticketcollection.InsertOne(t.ctx, ticket)
	if _, err := kf.Produce(ctx, "ticket", ticket); err != nil {
		return err
	}
	return err
}

func (t *TicketServiceImpl) GetTicket(pnr_number *string) (*models.Ticket, error){
	var ticket *models.Ticket
	fmt.Println("pnr_num: ", pnr_number)
	query := bson.D{bson.E{Key:"pnr_number", Value: pnr_number}}
	err := t.ticketcollection.FindOne(t.ctx, query).Decode(&ticket)
	return ticket, err
}

func (t *TicketServiceImpl) GetAll() ([]*models.Ticket, error){
	var tickets []*models.Ticket
	cursor, err := t.ticketcollection.Find(t.ctx, bson.D{{}})
	// fmt.Println("error1: ", err.Error())
	if err != nil{
		return nil, err
	}
	for cursor.Next(t.ctx){
		var ticket models.Ticket
		err := cursor.Decode(&ticket)
		if err != nil {
			return nil, err
		}
		// fmt.Println("error2: ", err.Error())
		tickets = append(tickets, &ticket)
	}
	if err := cursor.Err(); err != nil{
		// fmt.Println("error3: ", err.Error())
		return nil, err
	}

	cursor.Close(t.ctx)

	if len(tickets) == 0 {
		return nil, errors.New("documents not found")
	}
	return tickets, nil
}

func (t *TicketServiceImpl) UpdateTicket(ticket *models.Ticket) error{
	filter := bson.D{bson.E{Key:"pnr_number", Value: ticket.PNR_number}}
	update := bson.D{
		bson.E{
			Key:"$set", 
			Value: bson.D{
				bson.E{Key:"pnr_number", Value: ticket.PNR_number}, 
				bson.E{Key:"train_number", Value: ticket.Train_number}, 
				bson.E{Key:"seat_number", Value: ticket.Seat_number},
				bson.E{Key:"date_time", Value: ticket.Date_time},
				bson.E{Key:"passenger_name", Value: ticket.Passenger_name},
				bson.E{Key:"source", Value: ticket.Source},
				bson.E{Key:"destination", Value: ticket.Destination},
			}}}
	result,_ := t.ticketcollection.UpdateOne(t.ctx, filter, update)
	if result.MatchedCount != 1{
		return errors.New("no match found for update")
	}
	return nil
}

func (u *TicketServiceImpl) DeleteTicket(name *string) error{
	// filter := bson.D{bson.E{Key:"user_name", Value: name}}
	// result, _ = t.usercollection.DeleteOne(t.ctx, filter)
	// if result.DeletedCount != 1{
	// 	return errors.New("no match found for delete")
	// } 	
	return nil
}
