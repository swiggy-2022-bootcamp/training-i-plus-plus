// Code generated by mockery v2.10.2. DO NOT EDIT.

package mocks

import (
	domain "user/domain"
	errs "user/utils/errs"

	mock "github.com/stretchr/testify/mock"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// DeleteByEmail provides a mock function with given fields: _a0
func (_m *UserService) DeleteByEmail(_a0 string) (*domain.User, *errs.AppError) {
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

// FindByEmail provides a mock function with given fields: _a0
func (_m *UserService) FindByEmail(_a0 string) (*domain.User, *errs.AppError) {
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

// Login provides a mock function with given fields: _a0, _a1
func (_m *UserService) Login(_a0 string, _a1 string) (string, *errs.AppError) {
	ret := _m.Called(_a0, _a1)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(string, string) *errs.AppError); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}

// Register provides a mock function with given fields: _a0
func (_m *UserService) Register(_a0 domain.User) (*domain.User, *errs.AppError) {
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

// Update provides a mock function with given fields: _a0
func (_m *UserService) Update(_a0 domain.User) (*domain.User, *errs.AppError) {
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

// VerifyCredentials provides a mock function with given fields: _a0, _a1
func (_m *UserService) VerifyCredentials(_a0 string, _a1 string) (bool, *errs.AppError) {
	ret := _m.Called(_a0, _a1)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(string, string) *errs.AppError); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}

// VerifyToken provides a mock function with given fields: _a0, _a1, _a2
func (_m *UserService) VerifyToken(_a0 string, _a1 string, _a2 string) (bool, *errs.AppError) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string, string) bool); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(string, string, string) *errs.AppError); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}