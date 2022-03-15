package controller

import (
	"encoding/json"
	"net/http"
	db "ticket_reservation_system/database"
	user_model "ticket_reservation_system/model"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

var users []user_model.User
var userCollection *mongo.Collection = db.DatabaseConn().Database("maindb").Collection("users")

func AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user user_model.User
	json.NewDecoder(r.Body).Decode(&user)
	users = append(users, user)
	resp_string := "user created with username: " + user.UserName
	json.NewEncoder(w).Encode(resp_string)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputUserId := params["user_id"]
	for _, user := range users {
		if user.UserId == inputUserId {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputUserId := params["user_id"]
	for i, user := range users {
		if user.UserId == inputUserId {
			users = append(users[:i], users[i+1:]...)
			var updateUser user_model.User
			json.NewDecoder(r.Body).Decode(&updateUser)
			users = append(users, updateUser)
			json.NewEncoder(w).Encode(updateUser)
			return
		}
	}
}

func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputUserId := params["user_id"]
	for i, user := range users {
		if user.UserId == inputUserId {
			users = append(users[:i], users[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
}
