// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"

	mock "github.com/stretchr/testify/mock"
)

// HandlerUser is an autogenerated mock type for the HandlerUser type
type HandlerUser struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields:
func (_m *HandlerUser) CreateUser() gin.HandlerFunc {
	ret := _m.Called()

	var r0 gin.HandlerFunc
	if rf, ok := ret.Get(0).(func() gin.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gin.HandlerFunc)
		}
	}

	return r0
}

// DeleteUserID provides a mock function with given fields:
func (_m *HandlerUser) DeleteUserID() gin.HandlerFunc {
	ret := _m.Called()

	var r0 gin.HandlerFunc
	if rf, ok := ret.Get(0).(func() gin.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gin.HandlerFunc)
		}
	}

	return r0
}

// GetAllUsers provides a mock function with given fields:
func (_m *HandlerUser) GetAllUsers() gin.HandlerFunc {
	ret := _m.Called()

	var r0 gin.HandlerFunc
	if rf, ok := ret.Get(0).(func() gin.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gin.HandlerFunc)
		}
	}

	return r0
}

// GetUserID provides a mock function with given fields:
func (_m *HandlerUser) GetUserID() gin.HandlerFunc {
	ret := _m.Called()

	var r0 gin.HandlerFunc
	if rf, ok := ret.Get(0).(func() gin.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gin.HandlerFunc)
		}
	}

	return r0
}

// Login provides a mock function with given fields:
func (_m *HandlerUser) Login() gin.HandlerFunc {
	ret := _m.Called()

	var r0 gin.HandlerFunc
	if rf, ok := ret.Get(0).(func() gin.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gin.HandlerFunc)
		}
	}

	return r0
}

// UpdateUserID provides a mock function with given fields:
func (_m *HandlerUser) UpdateUserID() gin.HandlerFunc {
	ret := _m.Called()

	var r0 gin.HandlerFunc
	if rf, ok := ret.Get(0).(func() gin.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gin.HandlerFunc)
		}
	}

	return r0
}

type NewHandlerUserT interface {
	mock.TestingT
	Cleanup(func())
}

// NewHandlerUser creates a new instance of HandlerUser. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHandlerUser(t NewHandlerUserT) *HandlerUser {
	mock := &HandlerUser{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}