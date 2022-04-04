// Code generated by mockery v2.10.2. DO NOT EDIT.

package mocks

import (
	domain "github.com/swiggy-2022-bootcamp/training-i-plus-plus/ayan/mini-project/shopping-app/product/domain"
	errs "github.com/swiggy-2022-bootcamp/training-i-plus-plus/ayan/mini-project/shopping-app/product/utils/errs"

	mock "github.com/stretchr/testify/mock"
)

// ProductService is an autogenerated mock type for the ProductService type
type ProductService struct {
	mock.Mock
}

// DeleteById provides a mock function with given fields: _a0
func (_m *ProductService) DeleteById(_a0 string) (*domain.Product, *errs.AppError) {
	ret := _m.Called(_a0)

	var r0 *domain.Product
	if rf, ok := ret.Get(0).(func(string) *domain.Product); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Product)
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

// FindById provides a mock function with given fields: _a0
func (_m *ProductService) FindById(_a0 string) (*domain.Product, *errs.AppError) {
	ret := _m.Called(_a0)

	var r0 *domain.Product
	if rf, ok := ret.Get(0).(func(string) *domain.Product); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Product)
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

// Register provides a mock function with given fields: _a0
func (_m *ProductService) Register(_a0 domain.Product) (*domain.Product, *errs.AppError) {
	ret := _m.Called(_a0)

	var r0 *domain.Product
	if rf, ok := ret.Get(0).(func(domain.Product) *domain.Product); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Product)
		}
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(domain.Product) *errs.AppError); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *ProductService) Update(_a0 string, _a1 domain.Product) (*domain.Product, *errs.AppError) {
	ret := _m.Called(_a0, _a1)

	var r0 *domain.Product
	if rf, ok := ret.Get(0).(func(string, domain.Product) *domain.Product); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Product)
		}
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(string, domain.Product) *errs.AppError); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}
