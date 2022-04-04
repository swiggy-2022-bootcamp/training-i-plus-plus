// Code generated by mockery v2.10.2. DO NOT EDIT.

package mocks

import (
	domain "user/domain"
	errs "user/utils/errs"

	mock "github.com/stretchr/testify/mock"
)

// UserRepositoryDB is an autogenerated mock type for the UserRepositoryDB type
type UserRepositoryDB struct {
	mock.Mock
}

// DeleteUserByEmail provides a mock function with given fields: _a0
func (_m *UserRepositoryDB) DeleteUserByEmail(_a0 string) *errs.AppError {
	ret := _m.Called(_a0)

	var r0 *errs.AppError
	if rf, ok := ret.Get(0).(func(string) *errs.AppError); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errs.AppError)
		}
	}

	return r0
}

// FetchUserByEmail provides a mock function with given fields: _a0
func (_m *UserRepositoryDB) FetchUserByEmail(_a0 string) (*domain.User, *errs.AppError) {
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

// Save provides a mock function with given fields: _a0
func (_m *UserRepositoryDB) Save(_a0 domain.User) (*domain.User, *errs.AppError) {
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

// UpdateUser provides a mock function with given fields: _a0
func (_m *UserRepositoryDB) UpdateUser(_a0 domain.User) (*domain.User, *errs.AppError) {
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
