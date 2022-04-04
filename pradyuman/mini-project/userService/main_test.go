package main

import (
	"encoding/json"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var router = setupRouter()

func TestGetAuserValidUserId(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	UserId := "624a9f8ffbbe9e1b1e4882e5"
	req, _ := http.NewRequest("GET", "/user/"+UserId, nil)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	var reqBody map[string]interface{}
	json.Unmarshal([]byte(w.Body.String()), &reqBody)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "success", reqBody["message"])
}

func TestGetAuserInValidUserId(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	UserId := "6241a04313f320b1ecc10444"
	req, _ := http.NewRequest("GET", "/user/"+UserId, nil)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	var reqBody map[string]interface{}
	json.Unmarshal([]byte(w.Body.String()), &reqBody)
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "Invalid User Id", reqBody["message"])
}

func TestGetAlluserValid(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users", nil)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	var reqBody map[string]interface{}
	json.Unmarshal([]byte(w.Body.String()), &reqBody)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "success", reqBody["message"])
}

func TestDeleteAuserValidUserId(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	UserId := "624a9f8cfbbe9e1b1e4882e3"
	req, _ := http.NewRequest("DELETE", "/user/"+UserId, nil)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	var reqBody map[string]interface{}
	json.Unmarshal([]byte(w.Body.String()), &reqBody)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "User successfully deleted!", reqBody["message"])
}

func TestDeleteAuserInValidUserId(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	UserId := "624a9f7c1d96aa56f421e005"
	req, _ := http.NewRequest("DELETE", "/user/"+UserId, nil)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	var reqBody map[string]interface{}
	json.Unmarshal([]byte(w.Body.String()), &reqBody)
	assert.Equal(t, 404, w.Code)
	assert.Equal(t, "User with specified ID not found!", reqBody["message"])
}

// func TestSignUpInvalidUserInfo(t *testing.T) {
// 	router := setupRouter()

// 	var jsonStr = []byte(`{
// 		"name":"shilpi",
// 		"email":"sa33@iitbbs.ac.in",
// 		"password":"123456789"
// 	}`)

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonStr))
// 	req.Header.Set("Content-Type", "application/json")
// 	router.ServeHTTP(w, req)

// 	var reqBody map[string]interface{}
// 	json.Unmarshal([]byte(w.Body.String()),&reqBody)
// 	assert.Equal(t, 400, w.Code)
// 	assert.Equal(t, "error", reqBody["message"])
// }

// func TestLogin(t *testing.T) {

// 	var jsonStr = []byte(`{
// 		"email":"sa33@iitbbs.ac.in",
// 		"password":"123456789"
// 	}`)

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonStr))
// 	req.Header.Set("Content-Type", "application/json")
// 	router.ServeHTTP(w, req)
// 	var reqBody map[string]interface{}
// 	json.Unmarshal([]byte(w.Body.String()),&reqBody)
// 	assert.Equal(t, 201, w.Code)
// 	assert.Equal(t, "success", reqBody["message"])
// 	assert.NotNil(t,reqBody["jwt"])

// }

// func TestLoginPasswordMissing(t *testing.T) {

// 	var jsonStr = []byte(`{
// 		"email":"sa33@iitbbs.ac.in"
// 	}`)

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonStr))
// 	req.Header.Set("Content-Type", "application/json")
// 	router.ServeHTTP(w, req)
// 	var reqBody map[string]interface{}
// 	json.Unmarshal([]byte(w.Body.String()),&reqBody)
// 	assert.Equal(t, 400, w.Code)
// 	assert.Equal(t, "Login Info Missing", reqBody["message"])
// 	assert.Equal(t,"",reqBody["jwt"])
// }

// func TestLoginPasswordInvalid(t *testing.T) {

// 	var jsonStr = []byte(`{
// 		"password":"12345",
// 		"email":"sa33@iitbbs.ac.in"
// 	}`)

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonStr))
// 	req.Header.Set("Content-Type", "application/json")
// 	router.ServeHTTP(w, req)
// 	var reqBody map[string]interface{}
// 	json.Unmarshal([]byte(w.Body.String()),&reqBody)
// 	assert.Equal(t, 401, w.Code)
// 	assert.Equal(t, "error", reqBody["message"])
// 	assert.Equal(t,"",reqBody["jwt"])
// }

// func TestLoginEmailInvalid(t *testing.T) {

// 	var jsonStr = []byte(`{
// 		"password":"123456789",
// 		"email":"sa33@ibs.ac.in"
// 	}`)

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonStr))
// 	req.Header.Set("Content-Type", "application/json")
// 	router.ServeHTTP(w, req)
// 	var reqBody map[string]interface{}
// 	json.Unmarshal([]byte(w.Body.String()),&reqBody)
// 	assert.Equal(t, 400, w.Code)
// 	assert.Equal(t, "Invalid Email Id", reqBody["message"])
// 	assert.Equal(t,"",reqBody["jwt"])
// }
