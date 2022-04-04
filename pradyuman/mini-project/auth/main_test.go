package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var router = setupRouter()

func TestSignUpNewUser(t *testing.T) {
	

	var jsonStr = []byte(`{
		"name":"shilpi",
		"email":"sa33@iitbbs.ac.in",
		"password":"123456789",
		"phone":"9876543210",
		"role":"BUYER"
	}`)
	
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	var reqBody map[string]interface{}
	json.Unmarshal([]byte(w.Body.String()),&reqBody) 	
	assert.Equal(t, 201, w.Code)
	assert.Equal(t, "success", reqBody["message"])
}

func TestSignUpInvalidUserInfo(t *testing.T) {
	//router := setupRouter()

	var jsonStr = []byte(`{
		"name":"shilpi",
		"email":"sa33@iitbbs.ac.in",
		"password":"123456789"
	}`)
	
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	var reqBody map[string]interface{}
	json.Unmarshal([]byte(w.Body.String()),&reqBody) 	
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "error", reqBody["message"])
}

func TestLogin(t *testing.T) {
	
	var jsonStr = []byte(`{
		"email":"sa33@iitbbs.ac.in",
		"password":"123456789"
	}`)
	
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	var reqBody map[string]interface{}
	json.Unmarshal([]byte(w.Body.String()),&reqBody) 	
	assert.Equal(t, 201, w.Code)
	assert.Equal(t, "success", reqBody["message"])
	assert.NotNil(t,reqBody["jwt"])

}

func TestLoginPasswordMissing(t *testing.T) {
	
	var jsonStr = []byte(`{
		"email":"sa33@iitbbs.ac.in"
	}`)
	
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	var reqBody map[string]interface{}
	json.Unmarshal([]byte(w.Body.String()),&reqBody) 	
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "Login Info Missing", reqBody["message"])
	assert.Equal(t,"",reqBody["jwt"])	
}

func TestLoginPasswordInvalid(t *testing.T) {
	
	var jsonStr = []byte(`{
		"password":"12345",
		"email":"sa33@iitbbs.ac.in"
	}`)
	
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	var reqBody map[string]interface{}
	json.Unmarshal([]byte(w.Body.String()),&reqBody) 	
	assert.Equal(t, 401, w.Code)
	assert.Equal(t, "error", reqBody["message"])
	assert.Equal(t,"",reqBody["jwt"])
}

func TestLoginEmailInvalid(t *testing.T) {
	
	var jsonStr = []byte(`{
		"password":"123456789",
		"email":"sa33@ibs.ac.in"
	}`)
	
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	var reqBody map[string]interface{}
	json.Unmarshal([]byte(w.Body.String()),&reqBody) 	
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "Invalid Email Id", reqBody["message"])
	assert.Equal(t,"",reqBody["jwt"])
}