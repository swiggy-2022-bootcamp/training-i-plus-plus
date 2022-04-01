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
)
const (
    brokerAddress = "localhost:9092"
    topic         = "ticket"
)

var (
    traincontroller controllers.TrainController
)

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
        
        t := model.Ticket{}
		json.Unmarshal(msg.Value, &t)
        
        coll = main.Traincollection
        fmt.Println("message:: ", t.Train_number, coll)
    }
}
