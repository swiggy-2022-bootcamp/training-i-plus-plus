package domain

import (
	"alfred/external"
	"alfred/utils/errs"
	"strconv"
)

type CartService interface {
	AddToCart(int, map[string]int) *errs.AppError
	GetCart(int) (*Cart, *errs.AppError)
	CheckoutCart(int) ([]string, *errs.AppError)
}

type service struct {
	cartRepository CartRepository
	orderService   OrderService
}

func (s service) AddToCart(userId int, items map[string]int) *errs.AppError {
	return s.cartRepository.AddToCart(userId, items)
}

func (s service) GetCart(userId int) (*Cart, *errs.AppError) {
	return s.cartRepository.GetCart(userId)
}

func (s service) CheckoutCart(userId int) ([]string, *errs.AppError) {
	totalCartValue := float64(0)
	cart, err := s.cartRepository.GetCart(userId)
	if err != nil {
		return nil, err
	}
	itemsOutOfStock := make([]string, 0)
	flag := 0
	newItemQuantities := make(map[string]int)
	for itemId := range cart.Items {
		itemQuantityInCart := cart.Items[itemId]
		inventoryItem, err := external.GetItemByItemId(itemId)
		if err != nil {
			return nil, err
		}
		itemQuantityInInventory := inventoryItem.Quantity
		if itemQuantityInCart > itemQuantityInInventory {
			itemsOutOfStock = append(itemsOutOfStock, itemId)
			flag = 1
		} else {
			newItemQuantities[itemId] = itemQuantityInInventory - itemQuantityInCart
			totalCartValue = totalCartValue + float64(itemQuantityInCart)*inventoryItem.Price
		}
	}

	if flag == 0 {
		//TODO kafka call to update the total amount spent

		for k := range newItemQuantities {
			newQuantity := newItemQuantities[k]
			external.UpdateQuantity(k, newQuantity)
		}
		s.cartRepository.RemoveCart(userId)
		newOrder, _ := s.orderService.CreateOrder(userId, totalCartValue, cart.Items)

		res := []string{strconv.Itoa(newOrder.Id)}
		return res, nil
	} else {
		return itemsOutOfStock, errs.NewNotFoundError("Some items are out of stock / insufficient quantity available")
	}
}

func NewCartService(cartRepository CartRepository, orderService OrderService) CartService {
	return &service{
		cartRepository: cartRepository,
		orderService:   orderService,
	}
}
