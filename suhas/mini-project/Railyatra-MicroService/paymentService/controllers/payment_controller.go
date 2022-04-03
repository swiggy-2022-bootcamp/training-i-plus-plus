package controllers

import (
	"context"
	"fmt"
	"net"
	config "paymentService/config"
	log "paymentService/logger"
	"paymentService/models"
	"time"

	pb "paymentService/protobuf"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

var (
	collectionAuthName = "payment"
	collectionAuth     = new(mongo.Collection)
	errLog             = log.ErrorLogger.Println
)

type Server struct {
	pb.UnimplementedChargeServiceServer
}

func init() {
	var DB *mongo.Client = config.ConnectDB()
	collectionAuth = DB.Database("golangAPI").Collection(collectionAuthName)
}

type PaymentRepository struct{}

func (pr PaymentRepository) Insert(newCh models.Charge) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collectionAuth.InsertOne(ctx, &newCh)
	if err != nil {
		errLog(err)
	}
	return result.InsertedID, err
}

func SavePayment(ch *models.Charge) (err error) {
	pay := PaymentRepository{}
	_, err = pay.Insert(*ch)
	return err
}

//grpc

func (s *Server) Charge(ctx context.Context, in *pb.ChargeRequest) (*pb.ChargeResponse, error) {

	apiKey := config.EnvStripeSecret()
	fmt.Println(apiKey + "asdasd")
	// stripe.Key = apiKey
	// _, err := charge.New(&stripe.ChargeParams{
	// 	Amount:       stripe.Int64(int64(in.Amount)),
	// 	Currency:     stripe.String(string(stripe.CurrencyUSD)),
	// 	Description:  stripe.String("Ticket Bookging for" + in.Ticketid),
	// 	Source:       &stripe.SourceParams{Token: stripe.String("tok_visa")},
	// 	ReceiptEmail: stripe.String(in.Receiptemail)})

	stripe.Key = apiKey

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(in.Amount)),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		PaymentMethodTypes: []*string{
			stripe.String("card"),
		},
	}
	pi, err := paymentintent.New(params)
	fmt.Print(pi)
	if err != nil {
		errLog("Payment Unsuccesfull")
		return &pb.ChargeResponse{
			Confirmation: false,
			Message:      "Payment Unsuccesfull",
		}, err
	}
	newch := models.Charge{
		Amount:       int64(in.Amount),
		ReceiptEmail: in.Receiptemail,
		TicketID:     in.Ticketid,
	}
	err = SavePayment(&newch)
	if err != nil {
		errLog("Server Down !!!")
		return &pb.ChargeResponse{
			Confirmation: false,
			Message:      "Server Down !!!",
		}, err
	}

	return &pb.ChargeResponse{
		Confirmation: true,
		Message:      "Payment Succesfull",
	}, nil
}

func GrpcPaymentServer() error {
	gr := grpc.NewServer()
	lis, err := net.Listen("tcp", ":6012")
	if err != nil {
		errLog("Failed to listen: %v", err)
		fmt.Printf("Failed to listen: %v\n", err)
		return err
	}

	pb.RegisterChargeServiceServer(gr, &Server{})
	err = gr.Serve(lis)
	if err != nil {
		errLog("Failed to serve: %v", err)
		fmt.Printf("Failed to serve: %v\n", err)
		return err
	}
	return nil
}
