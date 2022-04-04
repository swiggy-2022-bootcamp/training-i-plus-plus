package app

import (
	"encoding/json"
	"fmt"
	"github.com/swiggy-2022-bootcamp/training-i-plus-plus/murtaza/mini-project/inventory/domain"
	"github.com/swiggy-2022-bootcamp/training-i-plus-plus/murtaza/mini-project/inventory/utils/errs"
	"github.com/swiggy-2022-bootcamp/training-i-plus-plus/murtaza/mini-project/inventory/utils/logger"
	"go.uber.org/zap"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ItemHandler struct {
	itemService domain.ItemService
}

type itemDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

type itemResponseDTO struct {
	Message string `json:"message"`
}

func (h ItemHandler) getAllUsers(c *gin.Context) {

}

// @Schemes
// @Description Fetches item details by itemId
// @Tags users
// @Param        itemId   path      int  true  "User Id"
// @Produce json
// @Success 200 {object} domain.Item
// @Failure      404  {object} errs.AppError
// @Failure      500  {object} errs.AppError
// @Router /items/{itemId} [get]
func (h ItemHandler) getItemByItemId(c *gin.Context) {
	params := c.Params
	param, _ := params.Get("itemId")
	itemId, err := strconv.ParseInt(param, 10, 0)
	if err != nil {
		logger.Error("Mandatory field itemId misisng in request params:")
		c.JSON(http.StatusBadRequest, errs.NewBadRequest("Mandatory field itemId missing in request params"))
		c.Abort()
		return
	}
	item, err2 := h.itemService.GetItemById(int(itemId))
	if err2 != nil {
		c.JSON(err2.Code, err2)
		c.Abort()
		return
	} else {
		logger.Info(fmt.Sprintf("Sending item details for itemId: %d", itemId))
		c.JSON(http.StatusOK, item)
	}
}

// @Schemes
// @Description Fetches item details by item name
// @Tags users
// @Param        name   query      string  true  "name"
// @Produce json
// @Success 200 {object} domain.Item
// @Failure      404  {object} errs.AppError
// @Failure      500  {object} errs.AppError
// @Router /items [get]
func (h ItemHandler) getItemByItemName(c *gin.Context) {
	name := c.Query("name")

	if name == "" {
		logger.Error("Mandatory field 'name' missing in request params:")
		c.JSON(http.StatusBadRequest, errs.NewBadRequest("Mandatory field 'name' missing in query params"))
		c.Abort()
		return
	}
	item, err2 := h.itemService.GetItemByName(name)
	if err2 != nil {
		c.JSON(err2.Code, err2)
		c.Abort()
		return
	} else {
		logger.Info(fmt.Sprintf("Sending item details for item name: %s", name))
		c.JSON(http.StatusOK, item)
	}
}

// @Schemes
// @Description Creates a new item
// @Tags users
// @Produce json
// @Accept json
// @Param        item  body      itemDTO  true  "create item"
// @Success 200 {object} domain.Item
// @Router /items [post]
func (h ItemHandler) createItem(c *gin.Context) {
	var newItem itemDTO
	err := json.NewDecoder(c.Request.Body).Decode(&newItem)
	if err != nil {
		customErr := errs.NewUnexpectedError("Unable to decode user payload")
		logger.Error(customErr.Message, zap.Error(err))
		c.JSON(http.StatusInternalServerError, customErr)
	} else {
		item, err := h.itemService.CreateItem(newItem.Name, newItem.Description, newItem.Quantity)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		} else {
			logger.Info(fmt.Sprintf(fmt.Sprintf("item created in DB with Id: %d", item.Id)))
			c.JSON(http.StatusCreated, item)
		}
	}
}

// @Schemes
// @Description Deletes item by itemId
// @Tags users
// @Param        itemId   path      int  true  "item ID"
// @Produce json
// @Success 200 {object} itemResponseDTO
// @Failure      500  {object} errs.AppError
// @Router /items/{itemId} [delete]
func (h ItemHandler) deleteItem(c *gin.Context) {
	params := c.Params
	val, err := params.Get("itemId")
	itemId, _ := strconv.Atoi(val)

	if err == false {
		logger.Error("Mandatory field 'itemId' missing in DELETE request")
		c.JSON(http.StatusBadRequest, gin.H{"message": "No item id given"})
	} else {
		err := h.itemService.DeleteItemById(itemId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		} else {
			logger.Error(fmt.Sprintf("Item with itemId: %d deleted successfully", itemId), zap.Int("itemId", itemId))
			itemResponse := itemResponseDTO{
				Message: fmt.Sprintf("itemId: %d deleted successfully", itemId),
			}
			c.JSON(http.StatusOK, itemResponse)
		}
	}
}

// @Schemes
// @Description Updates user by userId
// @Tags users
// @Param        itemId   path      int  true  "Item ID"
// @Param        item details   body      itemDTO true  "Item details"
// @Produce json
// @Success 200 {object} domain.Item
// @Failure      500  {object} errs.AppError
// @Router /items/{itemId} [put]
func (h ItemHandler) updateItem(c *gin.Context) {
	params := c.Params
	itemId, err := params.Get("itemId")

	if err == false {
		logger.Error("Mandatory field itemId missing in request")
		c.JSON(http.StatusBadRequest, "itemId missing in request")
	}

	var newItem itemDTO
	err2 := json.NewDecoder(c.Request.Body).Decode(&newItem)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, err2)
	} else {
		if err2 != nil {
			c.JSON(http.StatusInternalServerError, err2)
		}

		itemId, _ := strconv.ParseInt(itemId, 10, 0)
		item := domain.NewItem(newItem.Name, newItem.Description, newItem.Quantity)
		item.Id = int(itemId)
		updatedItem, err := h.itemService.UpdateItem(*item)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err2)
		} else {
			c.JSON(http.StatusOK, updatedItem)
		}
	}
}
