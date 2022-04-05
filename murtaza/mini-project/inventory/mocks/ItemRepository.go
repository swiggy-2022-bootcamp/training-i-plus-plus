// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/swiggy-2022-bootcamp/training-i-plus-plus/murtaza/mini-project/inventory/domain"
	errs "github.com/swiggy-2022-bootcamp/training-i-plus-plus/murtaza/mini-project/inventory/utils/errs"

	mock "github.com/stretchr/testify/mock"
)

// ItemRepository is an autogenerated mock type for the ItemRepository type
type ItemRepository struct {
	mock.Mock
}

// DeleteItemById provides a mock function with given fields: _a0
func (_m *ItemRepository) DeleteItemById(_a0 int) *errs.AppError {
	ret := _m.Called(_a0)

	var r0 *errs.AppError
	if rf, ok := ret.Get(0).(func(int) *errs.AppError); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errs.AppError)
		}
	}

	return r0
}

// FindItemById provides a mock function with given fields: _a0
func (_m *ItemRepository) FindItemById(_a0 int) (*domain.Item, *errs.AppError) {
	ret := _m.Called(_a0)

	var r0 *domain.Item
	if rf, ok := ret.Get(0).(func(int) *domain.Item); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Item)
		}
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(int) *errs.AppError); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}

// FindItemByName provides a mock function with given fields: _a0
func (_m *ItemRepository) FindItemByName(_a0 string) (*domain.Item, *errs.AppError) {
	ret := _m.Called(_a0)

	var r0 *domain.Item
	if rf, ok := ret.Get(0).(func(string) *domain.Item); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Item)
		}
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(string) *errs.AppError); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}

// InsertItem provides a mock function with given fields: _a0
func (_m *ItemRepository) InsertItem(_a0 domain.Item) (domain.Item, *errs.AppError) {
	ret := _m.Called(_a0)

	var r0 domain.Item
	if rf, ok := ret.Get(0).(func(domain.Item) domain.Item); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(domain.Item)
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(domain.Item) *errs.AppError); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}

// UpdateItem provides a mock function with given fields: _a0
func (_m *ItemRepository) UpdateItem(_a0 domain.Item) (*domain.Item, *errs.AppError) {
	ret := _m.Called(_a0)

	var r0 *domain.Item
	if rf, ok := ret.Get(0).(func(domain.Item) *domain.Item); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Item)
		}
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(domain.Item) *errs.AppError); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}

// UpdateItemQuantity provides a mock function with given fields: _a0, _a1
func (_m *ItemRepository) UpdateItemQuantity(_a0 int, _a1 int) *errs.AppError {
	ret := _m.Called(_a0, _a1)

	var r0 *errs.AppError
	if rf, ok := ret.Get(0).(func(int, int) *errs.AppError); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errs.AppError)
		}
	}

	return r0
}