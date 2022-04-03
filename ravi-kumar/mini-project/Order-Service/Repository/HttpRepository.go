package repository

import (
	"Order-Service/middleware"
	mockdata "Order-Service/model"
	"log"
	"net/http"
	"strconv"
)

type IHttpRepo interface {
	IsValidUser(userId string) bool
	UpdateProductQuantity(userId string, productIds []string, quantity int) (success bool, errorResponse *http.Response, errorProductIndex *int)
}

type HttpRepo struct {
}

func (httpRepo *HttpRepo) IsValidUser(userId string) bool {
	jwtToken, _ := middleware.GenerateJWT(userId, mockdata.Admin)
	url := "http://localhost:5004/users/" + userId

	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + jwtToken

	// Create a new request using http
	req, _ := http.NewRequest("GET", url, nil)

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode != http.StatusOK {
		return false
	}
	return true
}

func (httpRepo *HttpRepo) UpdateProductQuantity(userId string, productIds []string, quantity int) (success bool, errorResponse *http.Response, errorProductIndex *int) {
	jwtToken, _ := middleware.GenerateJWT(userId, mockdata.Admin)

	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + jwtToken

	for index, productId := range productIds {
		url := "http://localhost:5002/catalog/" + productId + "/" + strconv.Itoa(quantity)

		// Create a new request using http
		req, _ := http.NewRequest("POST", url, nil)

		// add authorization header to the req
		req.Header.Add("Authorization", bearer)

		// Send req using http Client
		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			log.Fatalln(err)
		}

		if resp.StatusCode != http.StatusOK {
			//roll back
			httpRepo.UpdateProductQuantity(userId, productIds[:index], 1)
			return false, resp, &index
		}
	}
	return true, nil, nil
}
