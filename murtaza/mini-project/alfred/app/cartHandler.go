package app

import (
	"alfred/domain"
	"alfred/external"
	"alfred/utils/errs"
	"alfred/utils/logger"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type CartHandler struct {
	cartService domain.CartService
}

type cartResponseDTO struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Data    []string `json:"data"`
}

func (ch CartHandler) addToCart(c *gin.Context) {
	userId := c.Request.Header.Get("userId")
	var items map[string]int
	err := json.NewDecoder(c.Request.Body).Decode(&items)
	if err != nil {
		customErr := errs.NewUnexpectedError("Unable to decode item payload")
		logger.Error(customErr.Message, zap.Error(err))
		response := cartResponseDTO{
			Code:    http.StatusInternalServerError,
			Message: customErr.Message,
		}
		c.JSON(http.StatusInternalServerError, response)
		c.Abort()
		return
	}
	id, _ := strconv.ParseInt(userId, 10, 0)
	myCart, err2 := ch.cartService.GetCart(int(id))
	for i := range items {
		var itemId = i
		item, err := external.GetItemByItemId(itemId)
		if err != nil {
			errMessage := fmt.Sprintf("Invalid Item Id: %s", itemId)
			logger.Error(errMessage, zap.String("itemId", itemId), zap.Error(err.Error()))
			response := cartResponseDTO{
				Code:    http.StatusBadRequest,
				Message: errMessage,
			}
			c.JSON(http.StatusBadRequest, response)
			c.Abort()
			return
		}

		totalQuantity := 0
		if err2 == nil {
			totalQuantity = myCart.Items[itemId] + items[itemId]
		} else {
			totalQuantity = items[itemId]
		}

		if item.Quantity < totalQuantity {
			errMessage := fmt.Sprintf("Insufficient Quantity for Item with item Id: %s", itemId)
			logger.Error(errMessage, zap.String("itemId", itemId))
			response := cartResponseDTO{
				Code:    http.StatusExpectationFailed,
				Message: errMessage,
			}
			c.JSON(http.StatusExpectationFailed, response)
			c.Abort()
			return
		}
	}
	err2 = ch.cartService.AddToCart(int(id), items)
	if err2 != nil {
		response := cartResponseDTO{
			Code:    http.StatusInternalServerError,
			Message: "something went wrong !",
		}
		logger.Error(err2.Message, zap.Error(err2.Error()))
		c.JSON(http.StatusInternalServerError, response)
		c.Abort()
		return
	}
	response := cartResponseDTO{
		Code:    http.StatusOK,
		Message: "Item(s) added To Cart successfully",
	}
	c.JSON(http.StatusOK, response)
}

func (ch CartHandler) checkoutCart(c *gin.Context) {
	userId := c.Request.Header.Get("userId")
	id, _ := strconv.ParseInt(userId, 10, 0)

	data, err := ch.cartService.CheckoutCart(int(id))
	if err != nil {
		if data != nil {
			response := cartResponseDTO{
				Code:    http.StatusExpectationFailed,
				Message: "some items are insufficient / out of stock",
				Data:    data,
			}
			c.JSON(response.Code, response)
			c.Abort()
			return
		}
		c.JSON(err.Code, err.Message)
		c.Abort()
		return
	}

	response := cartResponseDTO{
		Code:    http.StatusOK,
		Message: fmt.Sprintf("Order Placed Successfully with order Id: %s", data[0]),
		Data:    data,
	}

	c.JSON(http.StatusOK, response)
}
