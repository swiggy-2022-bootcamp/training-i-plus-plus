package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"search/internal/errors"
	model "search/internal/model"
	"search/util"
	"sync"

	"github.com/elastic/go-elasticsearch"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

type SearchProductService interface {
	ProcessRequest(ctx context.Context, searchedProduct string) ([]byte, *errors.ServerError)
}

var searchProductServiceStruct SearchProductService
var searchProductServiceOnce sync.Once

type searchProductService struct {
	config *util.RouterConfig
}

func InitSearchProductService(config *util.RouterConfig) SearchProductService {
	searchProductServiceOnce.Do(func() {
		searchProductServiceStruct = &searchProductService{
			config: config,
		}
	})

	return searchProductServiceStruct
}

func GetSearchProductService() SearchProductService {
	if searchProductServiceStruct == nil {
		panic("Add product service not initialised")
	}

	return searchProductServiceStruct
}

func (service *searchProductService) ProcessRequest(ctx context.Context, searchedProduct string) ([]byte, *errors.ServerError) {
	var goErr error

	productIds, err := service.search(searchedProduct)
	if err != nil {
		log.WithFields(logrus.Fields{
			"Error":            err,
			"Searched Product": searchedProduct,
		}).Error("an error occurred while searching for the product")
		return nil, err
	}

	if productIds == nil {
		log.WithField("Searched Product: ", searchedProduct).Info("No product found for given searched product item")
		return nil, nil
	}

	ids := model.ProductIds{Ids: productIds}
	var bodyBytes []byte

	reqBody, goErr := json.Marshal(ids)
	if goErr != nil {
		log.WithError(goErr).Error("an error occurred while marshalling the request body for Products service")
		return bodyBytes, &errors.InternalError
	}

	reqBodyBytes := bytes.NewBuffer(reqBody)

	resp, goErr := http.Post(service.config.WebServerConfig.ProductServiceUrl, "application/json", reqBodyBytes)
	if goErr != nil || resp.StatusCode != http.StatusOK {
		log.WithFields(logrus.Fields{
			"Error":      goErr,
			"StatusCode": resp.StatusCode,
		}).Error("an error occurred while calling to service: ", service.config.WebServerConfig.ProductServiceUrl)
		return bodyBytes, &errors.InternalError
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, goErr = io.ReadAll(resp.Body)
		if goErr != nil {
			log.WithField("Error: ", goErr).Error("an error occurred while reading the body bytes from response body")
			return bodyBytes, &errors.InternalError
		}
		bodyString := string(bodyBytes)
		log.Info(bodyString)
	}

	return bodyBytes, nil
}

func (service *searchProductService) search(searchedProduct string) ([]string, *errors.ServerError) {
	log.Info("inside search function ....")
	cfg := elasticsearch.Config{
		Addresses: []string{
			service.config.WebServerConfig.ElasticSearchUrl,
		},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	var buffer bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"combined_fields": map[string]interface{}{
				"query":  searchedProduct,
				"fields": []string{"title", "description"},
			},
		},
	}

	json.NewEncoder(&buffer).Encode(query)
	response, err := es.Search(es.Search.WithIndex("shopkart_data"), es.Search.WithBody(&buffer))
	if err != nil {
		log.WithField("Error: ", err).Error("an error occurred while calling to elastic search")
		return nil, nil
	}

	var result map[string]interface{}
	var productIds []string
	json.NewDecoder(response.Body).Decode(&result)
	for _, hit := range result["hits"].(map[string]interface{})["hits"].([]interface{}) {
		product := hit.(map[string]interface{})["_source"].(map[string]interface{})
		fmt.Println(product)
		fmt.Println(product["mongo_id"])
		productIds = append(productIds, product["mongo_id"].(string))
	}

	return productIds, nil
}
