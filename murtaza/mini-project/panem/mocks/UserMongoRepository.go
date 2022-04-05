// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	domain "panem/domain"
	errs "panem/utils/errs"

	mock "github.com/stretchr/testify/mock"
)

// UserMongoRepository is an autogenerated mock type for the UserMongoRepository type
type UserMongoRepository struct {
	mock.Mock
}

// DeleteUserByUserId provides a mock function with given fields: _a0
func (_m *UserMongoRepository) DeleteUserByUserId(_a0 int) *errs.AppError {
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

// FindUserById provides a mock function with given fields: _a0
func (_m *UserMongoRepository) FindUserById(_a0 int) (*domain.User, *errs.AppError) {
	ret := _m.Called(_a0)

	var r0 *domain.User
	if rf, ok := ret.Get(0).(func(int) *domain.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
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

// FindUserByUsername provides a mock function with given fields: _a0
func (_m *UserMongoRepository) FindUserByUsername(_a0 string) (*domain.User, *errs.AppError) {
	ret := _m.Called(_a0)

	var r0 *domain.User
	if rf, ok := ret.Get(0).(func(string) *domain.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
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

// InsertUser provides a mock function with given fields: _a0
func (_m *UserMongoRepository) InsertUser(_a0 domain.User) (domain.User, *errs.AppError) {
	ret := _m.Called(_a0)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(domain.User) domain.User); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(domain.User) *errs.AppError); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}

// UpdatePurchaseHistory provides a mock function with given fields: _a0, _a1, _a2
func (_m *UserMongoRepository) UpdatePurchaseHistory(_a0 int, _a1 int, _a2 float64) *errs.AppError {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *errs.AppError
	if rf, ok := ret.Get(0).(func(int, int, float64) *errs.AppError); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errs.AppError)
		}
	}

	return r0
}

// UpdateUser provides a mock function with given fields: _a0
func (_m *UserMongoRepository) UpdateUser(_a0 domain.User) (*domain.User, *errs.AppError) {
	ret := _m.Called(_a0)

	var r0 *domain.User
	if rf, ok := ret.Get(0).(func(domain.User) *domain.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(domain.User) *errs.AppError); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}
