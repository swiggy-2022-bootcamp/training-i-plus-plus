package external

import (
	"alfred/utils/errs"
	"alfred/utils/logger"
	"bytes"
	"encoding/json"
	"fmt"
	inventory "github.com/swiggy-2022-bootcamp/training-i-plus-plus/murtaza/mini-project/inventory/domain"
	"io"
	"net/http"
)

const InventoryApiUri = "http://localhost:8090/api/v1"

func GetItemByItemId(itemId string) (*inventory.Item, *errs.AppError) {
	resp, err := http.Get(fmt.Sprintf(InventoryApiUri+"/items/%s", itemId))
	if err != nil {
		logger.Fatal(err.Error())
	}
	var newItem inventory.Item
	err = json.NewDecoder(resp.Body).Decode(&newItem)
	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}
	return &newItem, nil
}

func UpdateQuantity(itemId string, quantity int) *errs.AppError {
	client := &http.Client{}
	url := fmt.Sprintf(InventoryApiUri+"/items/%s", itemId)
	payload, err := json.Marshal(map[string]interface{}{
		"quantity": quantity,
	})
	if err != nil {
		logger.Fatal(err.Error())
	}

	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(payload))
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
