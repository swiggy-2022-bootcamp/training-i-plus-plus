package domain

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"order/utils/errs"
	"order/utils/logger"
)

type OrderService interface {
	CreateOrder(Order) (*Order, *errs.AppError)
	FindById(string) (*Order, *errs.AppError)
	CheckAndCalcAmount([]OrderItem) (int, *errs.AppError)
	FindProductById(string) (*Product, *errs.AppError)
	UpdateProduct(string, Product) *errs.AppError
}

type DefaultOrderService struct {
	OrderDB OrderRepositoryDB
}

func NewOrderService(orderDB OrderRepositoryDB) OrderService {
	return &DefaultOrderService{
		OrderDB: orderDB,
	}
}

func (osvc DefaultOrderService) CreateOrder(order Order) (*Order, *errs.AppError) {

	amount, err := osvc.CheckAndCalcAmount(order.ItemList)
	if err != nil {
		return nil, err
	}
	order.Amount = amount

	u, err := osvc.OrderDB.Save(order)

	return u, err
}

func (osvc *DefaultOrderService) FindById(id string) (*Order, *errs.AppError) {

	order, err := osvc.OrderDB.FetchOrderById(id)
	return order, err
}

func (osvc *DefaultOrderService) CheckAndCalcAmount(itemList []OrderItem) (int, *errs.AppError) {

	amount := 0

	for _, item := range itemList {
		product, err := osvc.FindProductById(item.ProductId)
		if err != nil {
			return 0, err
		}
		if item.Quantity > product.Quantity {
			return 0, errs.NewValidationError(fmt.Sprintf("%s not available in required quantity", product.Name))
		}
		amount += product.Price * item.Quantity
		product.Quantity -= item.Quantity
		err = osvc.UpdateProduct(product.Id, *product)
		if err != nil {
			return 0, err
		}
	}
	return amount, nil
}

func (osvc *DefaultOrderService) FindProductById(productId string) (*Product, *errs.AppError) {

	resp, err := http.Get(fmt.Sprintf("http://localhost:8081/api/products/%s", productId))
	if err != nil {
		logger.Fatal(err.Error())
	}
	var product Product
	err = json.NewDecoder(resp.Body).Decode(&product)
	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}
	return &product, nil
}

func (osvc *DefaultOrderService) UpdateProduct(productId string, product Product) *errs.AppError {

	client := &http.Client{}
	url := fmt.Sprintf("http://localhost:8081/api/products/%s", productId)
	payload, err := json.Marshal(product)
	if err != nil {
		logger.Fatal(err.Error())
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		logger.Fatal(err.Error())
	}

	resp, err := client.Do(req)
	if err != nil {
		logger.Error(err.Error())
		return errs.NewUnexpectedError(err.Error())
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Error(err.Error())
		}
	}(resp.Body)

	return nil

}
