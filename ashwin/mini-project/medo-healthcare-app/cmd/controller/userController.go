package controller

import (
	"encoding/json"
	"fmt"
	"medo-healthcare-app/cmd/model"
	"medo-healthcare-app/pkg/authentication"
	"medo-healthcare-app/pkg/database"
	"net/http"
)

var activeSessionUser string

//ServeHome ..
func ServeHome(w http.ResponseWriter, r *http.Request) {
	if authentication.AuthenticateLogin(w, r) {
		if (GetUserType(w, r)) == "patient" {
			homePage := "./web/assets/patient/patientHomePage.html"
			http.ServeFile(w, r, homePage)
		} else if (GetUserType(w, r)) == "doctor" {
			homePage := "./web/assets/doctor/doctorHomePage.html"
			http.ServeFile(w, r, homePage)
		} else {
			homePage := "./web/assets/admin/adminHomePage.html"
			http.ServeFile(w, r, homePage)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		//json.NewEncoder(w).Encode("Unauthorized Access.\nPlease login and try again. :)")
		http.Error(w, "Unauthorized Access.\nPlease login and try again. :)", http.StatusUnauthorized)
	}
}

/* CRUD OPERATIONS */

//CreateUser ..
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var user model.CoreUserData
	_ = json.NewDecoder(r.Body).Decode(&user)
	database.InsertOne(user)
	json.NewEncoder(w).Encode(user)
}

//GetAllUsers ...
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	if authentication.AuthenticateLogin(w, r) {
		if (GetUserType(w, r)) == "masteradmin" {
			allUsersData := database.Find()
			json.NewEncoder(w).Encode(allUsersData)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			http.Error(w, "Unauthorized Access.\nPlease Try Logging Again with Administrator Access :)", http.StatusUnauthorized)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, "Unauthorized Access.\nPlease Login to access this URL ! :)", http.StatusUnauthorized)
	}
}

//GetOneUser ... Using Email Address
func GetOneUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if authentication.AuthenticateLogin(w, r) {
		if (GetUserType(w, r)) == "admin" || (GetUserType(w, r)) == "masteradmin" {
			var user model.CoreUserData
			_ = json.NewDecoder(r.Body).Decode(&user)
			result := database.FindOne(user.Email)
			json.NewEncoder(w).Encode(result)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			http.Error(w, "Unauthorized Access.\nPlease Try Logging Again with Administrator Access :)", http.StatusUnauthorized)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, "Unauthorized Access.\nPlease Login to access this URL ! :)", http.StatusUnauthorized)
	}
}

//UpdateOneUser ... To update a user using their username
func UpdateOneUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if authentication.AuthenticateLogin(w, r) {
		if (GetUserType(w, r)) == "admin" {
			var userToBeUpdated model.CoreUserData
			_ = json.NewDecoder(r.Body).Decode(&userToBeUpdated)
			result := database.UpdateOne(userToBeUpdated.Username, "email", "dkrv6666@gmail.com")
			json.NewEncoder(w).Encode(result)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, "Unauthorized Access.\nPlease Login to access this URL ! :)", http.StatusUnauthorized)
	}
}

//DeleteOneUser ..
func DeleteOneUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	if authentication.AuthenticateLogin(w, r) {
		if (GetUserType(w, r)) == "admin" {
			var user model.CoreUserData
			_ = json.NewDecoder(r.Body).Decode(&user)
			resultUser := database.DeleteOne(user.Email)
			json.NewEncoder(w).Encode(resultUser)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			http.Error(w, "Unauthorized Access.\nPlease Try Logging Again with Administrator Access :)", http.StatusUnauthorized)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, "Unauthorized Access.\nPlease Login to access this URL ! :)", http.StatusUnauthorized)
	}

}

//DeleteAllUsers ...
func DeleteAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	if authentication.AuthenticateLogin(w, r) {
		if (GetUserType(w, r)) == "masteradmin" {
			w.Write([]byte(GetUserType(w, r)))
			count := database.DeleteMany()
			json.NewEncoder(w).Encode(count)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			http.Error(w, "Unauthorized Access.\nPlease Try Logging Again with Administrator Access :)", http.StatusUnauthorized)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, "Unauthorized Access.\nPlease Login to access this URL ! :)", http.StatusUnauthorized)
	}
}

/* HELPER FUNCTIONS */

//GetUserType ... To determine whether the person is Doctor / Patient.
func GetUserType(w http.ResponseWriter, r *http.Request) string {
	resultUser := database.FindOne(authentication.GetUsernameFromToken(w, r))
	fmt.Println("Welcome ", resultUser.UserType, " ", resultUser.Name)
	activeSessionUser = resultUser.Username
	return resultUser.UserType
}

// //GetUsernameFromTokenController ...
// func GetUsernameFromTokenController(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println(authentication.GetUsernameFromToken(w, r))
// }
