package domain_test

import (
	"github.com/stretchr/testify/mock"
	"github.com/swiggy-2022-bootcamp/training-i-plus-plus/murtaza/mini-project/inventory/domain"
	"github.com/swiggy-2022-bootcamp/training-i-plus-plus/murtaza/mini-project/inventory/mocks"
	"github.com/swiggy-2022-bootcamp/training-i-plus-plus/murtaza/mini-project/inventory/utils/errs"
	"testing"

	"github.com/stretchr/testify/assert"
)

var mockItemRepo = mocks.ItemRepository{}
var itemService = domain.NewItemService(&mockItemRepo)

func TestShouldReturnNewItemService(t *testing.T) {
	userService := domain.NewItemService(nil)
	assert.NotNil(t, userService)
}

func TestShouldReturnTrueIfItemIsOutOfStock(t *testing.T) {
	name := "Cornflakes"
	description := "breakfast cereal 1kg"
	quantity := 0
	price := 150.3

	newItem := domain.NewItem(name, description, quantity, price)
	mockItemRepo.On("FindItemById", 1).Return(newItem, nil)
	actualResponse := itemService.IsItemOutOfStock(1)

	assert.True(t, actualResponse)
}

func TestShouldCreateNewItem(t *testing.T) {

	name := "Cornflakes"
	description := "breakfast cereal 1kg"
	quantity := 200
	price := 150.3

	newItem := domain.NewItem(name, description, quantity, price)

	mockItemRepo.On("InsertItem", mock.Anything).Return(*newItem, nil)
	itemService.CreateItem(name, description, quantity, price)
	mockItemRepo.AssertNumberOfCalls(t, "InsertItem", 1)
}

func TestShouldDeleteItemByItemId(t *testing.T) {
	itemId := 1
	mockItemRepo.On("DeleteItemById", itemId).Return(nil)

	var err = itemService.DeleteItemById(itemId)
	assert.Nil(t, err)
}

func TestShouldGetItemByItemId(t *testing.T) {
	itemId := 1
	name := "Cornflakes"
	description := "breakfast cereal 1kg"
	quantity := 0
	price := 150.3

	newItem := domain.NewItem(name, description, quantity, price)

	mockItemRepo.On("FindItemById", itemId).Return(newItem, nil)
	resItem, _ := itemService.GetItemById(itemId)

	assert.Equal(t, name, resItem.Name)
	assert.Equal(t, description, resItem.Description)
	assert.Equal(t, quantity, resItem.Quantity)
	assert.Equal(t, price, resItem.Price)
}

func TestShouldGetItemByItemName(t *testing.T) {
	name := "Cornflakes"
	description := "breakfast cereal 1kg"
	quantity := 200
	price := 150.3

	newItem := domain.NewItem(name, description, quantity, price)

	mockItemRepo.On("FindItemByName", name).Return(newItem, nil)
	resItem, _ := itemService.GetItemByName(name)

	assert.Equal(t, name, resItem.Name)
	assert.Equal(t, description, resItem.Description)
	assert.Equal(t, quantity, resItem.Quantity)
	assert.Equal(t, price, resItem.Price)
}

func TestShouldNotDeleteItemByItemIdUponInvalidItemId(t *testing.T) {
	itemId := -99
	errMessage := "some error"
	mockItemRepo.On("DeleteItemById", itemId).Return(errs.NewUnexpectedError(errMessage))

	err := itemService.DeleteItemById(itemId)
	assert.Error(t, err.Error(), errMessage)
}

func TestShouldUpdateItem(t *testing.T) {
	name := "Cornflakes"
	description := "breakfast cereal 1kg"
	quantity := 200
	price := 150.3

	newItem := domain.NewItem(name, description, quantity, price)
	mockItemRepo.On("UpdateItem", *newItem).Return(newItem, nil)
	updatedItem, _ := itemService.UpdateItem(*newItem)

	assert.Equal(t, newItem.Name, updatedItem.Name)
	assert.Equal(t, newItem.Description, updatedItem.Description)
	assert.Equal(t, newItem.Quantity, updatedItem.Quantity)
	assert.Equal(t, newItem.Price, updatedItem.Price)
}

func TestShouldUpdateItemQuantity(t *testing.T) {
	mockItemRepo.On("UpdateItemQuantity", 1, 100).Return(nil)
	err := itemService.UpdateItemQuantity(1, 100)
	assert.Nil(t, err)
}
