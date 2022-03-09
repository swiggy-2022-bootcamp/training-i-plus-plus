package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"mini-project/internal/mocks"
	"mini-project/internal/modals"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Users)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
			log.Fatalln(err)
	}

	user := &modals.User{}
	err = json.Unmarshal([]byte(body), user)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	user.Id = uuid.New()
	mocks.Users = append(mocks.Users, *user)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}


func GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	for _, user := range mocks.Users {
			if user.Id.String() == userId {
					w.Header().Add("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)

					json.NewEncoder(w).Encode(user)
					break
			}
	}
}

func UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"] 

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
			log.Fatalln(err)
	}


	updatedUser := &modals.User{}
	err = json.Unmarshal([]byte(body), updatedUser)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	for index, user := range mocks.Users {
			if user.Id.String() == userId {
					updatedUser.Id = user.Id
					mocks.Users[index] = *updatedUser
					w.Header().Add("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)

					json.NewEncoder(w).Encode("Updated User Details")
					break
			}
	}
}


func DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	for index, user := range mocks.Users {
			if user.Id.String() == userId {
					mocks.Users = append(mocks.Users[:index], mocks.Users[index+1:]...)

					w.Header().Add("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					json.NewEncoder(w).Encode("Deleted User Successfully")
					break
			}
	}
}