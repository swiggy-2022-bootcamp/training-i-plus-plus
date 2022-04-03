package kafkaConsumer

import (
    "fmt"
    "log"
    "os"
    // "strconv"
    // "time"
    "context"
    "encoding/json"
    "github.com/segmentio/kafka-go"
    model "github.com/swastiksahoo153/MicroserviceKafka/TrainModule/models"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
    "github.com/swastiksahoo153/MicroserviceKafka/TrainModule/database"
)
var (
    brokerAddress = "localhost:9092"
    topic         = "ticket"
    Traincollection	*mongo.Collection
	mongoclient 	*mongo.Client
    err         	error
)

func getIndex(s []int, num int) int {
    for i, n := range s{
        if n == num {
            return i
        }
    }
    return -1
}

func bookTicket(ctx context.Context, ticket model.Ticket) {

    mongoclient = database.GetDatabase(ctx)

    Traincollection = mongoclient.Database("traindb").Collection("trains")
    
    fmt.Println("ticket train no:: ", ticket.Train_number)
    train := model.Train{}
    query := bson.D{bson.E{Key:"train_number", Value: ticket.Train_number}}
    err = Traincollection.FindOne(ctx, query).Decode(&train)
    fmt.Println("train:: ", train)

    availableSeats := train.Seats_available

    index := getIndex(availableSeats, ticket.Seat_number)

    if index != -1 {
        train.Seats_available = append(availableSeats[:index], availableSeats[index+1:]...)
        fmt.Println("train: ", train)
        //modify in db
        filter := bson.D{bson.E{Key:"train_number", Value: train.Train_number}}
        update := bson.D{
            bson.E{
                Key:"$set", 
                Value: bson.D{
                    bson.E{Key:"train_number", Value: train.Train_number}, 
                    bson.E{Key:"train_name", Value: train.Train_name}, 
                    bson.E{Key:"source", Value: train.Source}, 
                    bson.E{Key:"destination", Value: train.Destination}, 
                    bson.E{Key:"seats_available", Value: train.Seats_available},
                    bson.E{Key:"total_seats", Value: train.Total_seats},
                }}}
    
        Traincollection.UpdateOne(ctx, filter, update)

    } else{
        //return ticket not available
        fmt.Println("not available!")
    }
}

func Consume(ctx context.Context) {
    l := log.New(os.Stdout, "kafka reader: ", 0)
    r := kafka.NewReader(kafka.ReaderConfig{
        Brokers: []string{brokerAddress},
        Topic:   topic,
        GroupID: "my-group",
        Logger: l,
    })
    for {
        msg, err := r.ReadMessage(ctx)
        if err != nil {
            panic("could not read message " + err.Error())
        }
        
        ticket := model.Ticket{}
		json.Unmarshal(msg.Value, &ticket)

        bookTicket(ctx, ticket)
    }
}
