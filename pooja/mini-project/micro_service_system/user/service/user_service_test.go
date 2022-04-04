package service

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestSignupFunctionality(t *testing.T) {
	apiUrl := "http://localhost:6001/user/login"
	req, _ := http.NewRequest("POST", apiUrl, nil)
	q := req.URL.Query()
	q.Add("username", "nana")
	q.Add("password", "pass")
	req.URL.RawQuery = q.Encode()
	fmt.Print(req.URL.String())
	req.Close = true
	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	fmt.Println("resp", resp.StatusCode)
	expectedStatus := http.StatusOK
	if resp.StatusCode != expectedStatus {
		t.Errorf("Expected %v but got %v ", expectedStatus, resp.Status)
	}

}
