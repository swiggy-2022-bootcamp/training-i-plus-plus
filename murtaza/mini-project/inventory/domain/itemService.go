package domain

import "github.com/swiggy-2022-bootcamp/training-i-plus-plus/murtaza/mini-project/inventory/utils/errs"

type ItemService interface {
	CreateItem(string, string, int) (Item, *errs.AppError)
	GetItemById(int) (*Item, *errs.AppError)
	GetItemByName(string) (*Item, *errs.AppError)
	DeleteItemById(int) *errs.AppError
	UpdateItem(Item) (*Item, *errs.AppError)
	IsItemOutOfStock(int) bool
}

type service struct {
	itemRepository ItemRepository
}

func (s service) CreateItem(name, description string, quantity int) (Item, *errs.AppError) {
	item := NewItem(name, description, quantity)
	persistedItem, err := s.itemRepository.InsertItem(*item)
	if err != nil {
		return Item{}, err
	}
	return persistedItem, nil
}

func (s service) GetItemById(itemId int) (*Item, *errs.AppError) {
	res, err := s.itemRepository.FindItemById(itemId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s service) GetItemByName(name string) (*Item, *errs.AppError) {
	res, err := s.itemRepository.FindItemByName(name)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s service) DeleteItemById(itemId int) *errs.AppError {
	err := s.itemRepository.DeleteItemById(itemId)
	if err != nil {
		return err
	}
	return nil
}

func (s service) UpdateItem(item Item) (*Item, *errs.AppError) {
	res, err := s.itemRepository.UpdateItem(item)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s service) IsItemOutOfStock(itemId int) bool {
	item, err := s.itemRepository.FindItemById(itemId)
	if err != nil {
		return true
	}
	return item.Quantity > 0
}

func NewItemService(itemRepository ItemRepository) ItemService {
	return &service{
		itemRepository: itemRepository,
	}
}
