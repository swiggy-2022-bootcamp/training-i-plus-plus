package controllers

import (
	"adminService/kafka"
	log "adminService/logger"
	"adminService/models"
	"adminService/repository"
	"adminService/responses"
	"context"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	pb "adminService/protobuf"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
)

//var trainCollection *mongo.Collection = config.GetCollection(config.DB, "trains")
//var availticketCollection *mongo.Collection = config.GetCollection(config.DB, "availtickets")
var avalidate = validator.New()

const layout = "Jan 2, 2006 at 3:04pm (MST)"

var (
	adminrepo       repository.AdminRepository
	trainrepo       repository.TrainRepository
	availticketrepo repository.AvailTicketRepository
	c               pb.AuthenticationServiceClient
	address         = "localhost:6010"
)

var (
	errLog  = log.ErrorLogger.Println
	infoLog = log.ErrorLogger.Println
)

type Server struct {
	pb.UnimplementedAvailTicketServiceServer
}

func init() {
	go kafka.Consume_booked_ticket_for_avail()

	//configure grpc client
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		errLog("Error while making connection, %v", err)
		fmt.Printf("Error while making connection, %v\n", err)
	}

	// Create a client instance
	c = pb.NewAuthenticationServiceClient(conn)
}

func respondWithError(c1 *gin.Context, code int, message interface{}) {
	c1.AbortWithStatusJSON(code, gin.H{"error": message})
}

func CheckAuthorized(group string) gin.HandlerFunc {
	return func(co *gin.Context) {
		bearToken := co.Request.Header.Get("Authorization")
		//normally Authorization the_token_xxx
		strArr := strings.Split(bearToken, " ")
		if len(strArr) != 2 {
			respondWithError(co, 401, "No bearer token")
			return
		}
		resp, err := c.Authenticate(
			context.Background(),
			&pb.AuthenticateRequest{
				Group: group,
				Token: strArr[1],
			})
		if err != nil || !resp.Confirmation {
			respondWithError(co, 401, err)
			return
		}
		infoLog("Autoriztion " + resp.Message)
		co.Next()
	}
}

func CreateAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var admin models.Admin
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&admin); err != nil {
			c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := avalidate.Struct(&admin); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newAdmin := models.Admin{
			Name:  admin.Name,
			Email: admin.Email,
		}

		result, err := adminrepo.Insert(newAdmin)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.AdminResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		adminId := c.Param("adminid")
		var admin models.Admin
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(adminId)

		admin, err := adminrepo.Read(objId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.AdminResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": admin}})
	}
}

func EditAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		adminId := c.Param("adminid")
		var admin models.Admin
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(adminId)

		//validate the request body
		if err := c.BindJSON(&admin); err != nil {
			c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := avalidate.Struct(&admin); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		result, err := adminrepo.Update(admin, objId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.AdminResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"updated id": result}})
	}
}

func DeleteAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		adminId := c.Param("adminid")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(adminId)

		resultcount, err := adminrepo.Delete(objId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if resultcount.(int) < 1 {
			c.JSON(http.StatusNotFound,
				responses.AdminResponse{Status: http.StatusNotFound, Message: err.Error(), Data: map[string]interface{}{"data": "Admin with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.AdminResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Admin successfully deleted!"}},
		)
	}
}

func GetAllAdmins() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var admins []models.Admin
		defer cancel()

		admins, err := adminrepo.ReadAll()

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK,
			responses.AdminResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": admins}},
		)
	}
}

func CreateTrain() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var train models.Train
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&train); err != nil {
			c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := avalidate.Struct(&train); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newTrain := models.Train{
			Station1: train.Station1,
			Station2: train.Station2,
		}

		result, err := trainrepo.Insert(newTrain)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		go kafka.Produce_train(newTrain)
		c.JSON(http.StatusCreated, responses.TrainResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetTrain() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		trainId := c.Param("trainid")
		var train models.Train
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(trainId)

		train, err := trainrepo.Read(objId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.TrainResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.TrainResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": train}})
	}
}

func EditTrain() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		trainId := c.Param("trainid")
		var train models.Train
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(trainId)

		//validate the request body
		if err := c.BindJSON(&train); err != nil {
			c.JSON(http.StatusBadRequest, responses.TrainResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := avalidate.Struct(&train); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.TrainResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		updatedTrainid, err := trainrepo.Update(train, objId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.TrainResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.TrainResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"id of the updated item": updatedTrainid}})
	}
}

func DeleteTrain() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		trainId := c.Param("trainid")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(trainId)

		resultCount, err := trainrepo.Delete(objId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.TrainResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if resultCount.(int) < 1 {
			c.JSON(http.StatusInternalServerError, responses.TrainResponse{Status: http.StatusInternalServerError, Message: "error id not found", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK,
			responses.TrainResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Train successfully deleted!"}},
		)
	}
}

func GetAllTrains() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var trains []models.Train
		defer cancel()

		trains, err := trainrepo.ReadAll()

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.TrainResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK,
			responses.TrainResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": trains}},
		)
	}
}

func CreateAvailTicket() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var availticket models.AvailTicket
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&availticket); err != nil {
			c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := avalidate.Struct(&availticket); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		_, err1 := time.Parse(layout, availticket.Departure_time)
		_, err := time.Parse(layout, availticket.Arrival_time)

		if err != nil || err1 != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "time not in correct format", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		newAvailTicket := models.AvailTicket{
			Train_id:       availticket.Train_id,
			Capacity:       availticket.Capacity,
			Price:          availticket.Price,
			Departure:      availticket.Departure,
			Arrival:        availticket.Arrival,
			Departure_time: availticket.Departure_time,
			Arrival_time:   availticket.Arrival_time,
		}

		result, err := availticketrepo.Insert(newAvailTicket)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		go kafka.Produce_avail_ticket(newAvailTicket)
		c.JSON(http.StatusCreated, responses.AvailTicketResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetAvailTicket() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		availticketId := c.Param("availticketid")
		var availticket models.AvailTicket
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(availticketId)

		//err := availticketCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&availticket)
		availticket, err := availticketrepo.Read(objId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AvailTicketResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.AvailTicketResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": availticket}})
	}
}

func EditAvailTicket() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		availticketId := c.Param("availticketid")
		var availticket models.AvailTicket
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(availticketId)

		//validate the request body
		if err := c.BindJSON(&availticket); err != nil {
			c.JSON(http.StatusBadRequest, responses.AvailTicketResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := avalidate.Struct(&availticket); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.AvailTicketResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		udpatedid, err := availticketrepo.Update(availticket, objId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AvailTicketResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.AvailTicketResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"updated id": udpatedid}})
	}
}

func DeleteAvailTicket() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		availticketId := c.Param("availticketid")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(availticketId)

		resultcount, err := availticketrepo.Delete(objId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AvailTicketResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if resultcount.(int) < 1 {
			c.JSON(http.StatusNotFound,
				responses.AvailTicketResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "AvailTicket with specified ID not found!"}},
			)
			return
		}
		c.JSON(http.StatusOK,
			responses.AvailTicketResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "AvailTicket successfully deleted!"}},
		)
	}
}

func GetAllAvailTickets() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var availtickets []models.AvailTicket
		defer cancel()

		availtickets, err := availticketrepo.ReadAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AvailTicketResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK,
			responses.AvailTicketResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": availtickets}},
		)
	}
}

//grpc

func (s *Server) GetTicketConfirmation(ctx context.Context, in *pb.AvailTicketRequest) (*pb.AvailTicketResponse, error) {
	objId, err := primitive.ObjectIDFromHex(in.TrainId)
	if err != nil {
		errLog("Incorrect format for train id")
		return &pb.AvailTicketResponse{
			Station1:      "",
			Station2:      "",
			ArrivalTime:   "",
			DepartureTime: "",
			Message:       1,
		}, err
	}
	avtick, err := availticketrepo.ReadTrainId(objId)
	if err != nil {
		errLog("Incorrect train id")
		return &pb.AvailTicketResponse{
			Station1:      "",
			Station2:      "",
			ArrivalTime:   "",
			DepartureTime: "",
			Message:       1,
		}, err
	}
	if uint32(avtick.Capacity) < in.NumOfTickets {
		errLog("Not enough tickets")
		return &pb.AvailTicketResponse{
			Station1:      "",
			Station2:      "",
			ArrivalTime:   "",
			DepartureTime: "",
			Message:       2,
		}, fmt.Errorf("not enough tickets")
	}

	return &pb.AvailTicketResponse{
		Station1:      avtick.Departure,
		Station2:      avtick.Arrival,
		ArrivalTime:   avtick.Arrival_time,
		DepartureTime: avtick.Departure_time,
		Message:       0,
	}, nil
}

func StartAdmingrpc() error {
	gr := grpc.NewServer()
	lis, err := net.Listen("tcp", ":6011")
	if err != nil {
		errLog("Failed to listen: %v", err)
		fmt.Printf("Failed to listen: %v\n", err)
		return err
	}

	pb.RegisterAvailTicketServiceServer(gr, &Server{})
	err = gr.Serve(lis)
	if err != nil {
		errLog("Failed to serve: %v", err)
		fmt.Printf("Failed to serve: %v\n", err)
		return err
	}
	return nil
}
