package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Server struct {
	*mux.Router
	*dynamodb.Client
}

func NewServer() *Server {
	s := &Server{
		Router: mux.NewRouter(),
		Client: CreateLocalClient(),
	}
	s.routes()
	return s
}

func CreateLocalClient() *dynamodb.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("mumbai"),
		config.WithEndpointResolver(aws.EndpointResolverFunc(
			func(service, region string) (aws.Endpoint, error) {
				return aws.Endpoint{URL: "http://localhost:8042"}, nil
			})),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID: "fake", SecretAccessKey: "fake",
			},
		}),
	)
	if err != nil {
		panic(err)
	}

	return dynamodb.NewFromConfig(cfg)
}

func (S *Server) routes() {
	S.Router.HandleFunc("/user/add", S.AddUserRoute()).Methods("POST")
	S.Router.HandleFunc("/user/read", S.ReadUserRoute()).Methods("POST")
	S.Router.HandleFunc("/user/update", S.UpdateUserRoute()).Methods("POST")
	S.Router.HandleFunc("/user/delete", S.DeleteUserRoute()).Methods("POST")
}

func (S *Server) AddUserRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var i User
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		i.UUID = uuid.New().String()
		AddUser(S.Client, i)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(i); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (S *Server) ReadUserRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//idStr := mux.Vars(r)["UUID"]
		type request struct {
			UUID string
		}
		var rq request
		fmt.Println(r.Body)
		if err := json.NewDecoder(r.Body).Decode(&rq); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(rq)

		// _, err := uuid.Parse(idStr)
		// if err != nil {
		// 	fmt.Println(err)
		// 	http.Error(w, err.Error(), http.StatusBadRequest)
		// 	return
		// }
		usr, err := ReadUser(S.Client, rq.UUID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNoContent)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(usr); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (S *Server) UpdateUserRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// idStr := mux.Vars(r)["UUID"]
		// _, err := uuid.Parse(idStr)
		// namestr := mux.Vars(r)["Name"]
		type requestUpdate struct {
			UUID  string
			Email string
		}
		var rqu requestUpdate
		fmt.Println(r.Body)
		if err := json.NewDecoder(r.Body).Decode(&rqu); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(rqu)

		err := UpdateUser(S.Client, rqu.UUID, rqu.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNoContent)
			return
		}
		type updateresp struct {
			Msg string
		}
		rsp := updateresp{
			"User Updated Succesfully",
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(rsp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (S *Server) DeleteUserRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type request struct {
			UUID string
		}
		var rq request
		fmt.Println(r.Body)
		if err := json.NewDecoder(r.Body).Decode(&rq); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(rq)

		err := DeleteUser(S.Client, rq.UUID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNoContent)
			return
		}
		type deleteresp struct {
			Msg string
		}
		rs := deleteresp{
			"User Deleted Succesfully",
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(rs); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
