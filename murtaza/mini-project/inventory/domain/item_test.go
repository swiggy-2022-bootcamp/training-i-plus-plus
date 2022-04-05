package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnNewItem(t *testing.T) {
	name := "Cornflakes"
	description := "breakfast cereal 1kg"
	quantity := 200
	price := 150.3

	newItem := NewItem(name, description, quantity, price)

	assert.Equal(t, name, newItem.Name)
	assert.Equal(t, description, newItem.Description)
	assert.Equal(t, quantity, newItem.Quantity)
	assert.Equal(t, price, newItem.Price)
}
