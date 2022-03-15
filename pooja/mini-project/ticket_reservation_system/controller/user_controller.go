package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	db "ticket_reservation_system/database"
	user_model "ticket_reservation_system/model"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var users []user_model.User
var userCollection *mongo.Collection = db.DatabaseConn().Database("maindb").Collection("users")

func AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user user_model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Println("error in decoding user payload with error ", err)
	}
	insertResult, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		fmt.Println("error in inserting the data to db ", err)
	}
	json.NewEncoder(w).Encode(insertResult.InsertedID)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var results []primitive.M
	cur, err := userCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		fmt.Println(err)
	}
	for cur.Next(context.TODO()) {
		var elem primitive.M
		if err := cur.Decode(&elem); err != nil {
			fmt.Println(err)
		}
		results = append(results, elem)
	}
	cur.Close(context.TODO())
	json.NewEncoder(w).Encode(results)
}

func GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user user_model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Print(err)
	}
	var result primitive.M //  an unordered representation of a BSON document which is a Map
	if err := userCollection.FindOne(context.TODO(), bson.D{{"username", user.UserName}}).Decode(&result); err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(result) // returns a Map containing document
}

func UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	type updateBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var body updateBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		fmt.Print(err)
	}
	filter := bson.D{{"username", body.Username}}
	after := options.After
	returnOpt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	update := bson.D{{"$set", bson.D{{"password", body.Password}}}}
	updateResult := userCollection.FindOneAndUpdate(context.TODO(), filter, update, &returnOpt)
	var result primitive.M
	_ = updateResult.Decode(&result)
	json.NewEncoder(w).Encode(result)
}

func DeleteUserByUsername(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	username := mux.Vars(r)["username"]
	opts := options.Delete().SetCollation(&options.Collation{})
	if _, err := userCollection.DeleteOne(context.TODO(), bson.D{{"username", username}}, opts); err != nil {
		fmt.Print(err)
	}
	json.NewEncoder(w).Encode(http.StatusNoContent)
}
