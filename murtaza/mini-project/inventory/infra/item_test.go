package infra

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldCreateNewUserPersistedEntity(t *testing.T) {

	name := "item-name"
	description := "item-description"
	quantity := 200
	price := 90.99

	persistedItem := NewItemModel(name, description, quantity, price)

	assert.Equal(t, name, persistedItem.Name)
	assert.Equal(t, description, persistedItem.Description)
	assert.Equal(t, quantity, persistedItem.Quantity)
	assert.Equal(t, price, persistedItem.Price)
	assert.NotNil(t, persistedItem.CreatedAt)
	assert.NotNil(t, persistedItem.UpdatedAt)
}

func TestShouldConvertUserPersistedEntityToDomainEntity(t *testing.T) {
	name := "item-name"
	description := "item-description"
	quantity := 200
	price := 90.99

	persistedItem := NewItemModel(name, description, quantity, price)

	domainItem := persistedItem.toDomainEntity()

	assert.Equal(t, persistedItem.Name, domainItem.Name)
	assert.Equal(t, persistedItem.Description, domainItem.Description)
	assert.Equal(t, persistedItem.Quantity, domainItem.Quantity)
	assert.Equal(t, persistedItem.Price, domainItem.Price)
}
