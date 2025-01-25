// Code generated by mockery v2.51.1. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/astronely/financial-helper_microservices/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, info, password
func (_m *UserRepository) Create(ctx context.Context, info *model.UserInfo, password string) (int64, string, error) {
	ret := _m.Called(ctx, info, password)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 int64
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.UserInfo, string) (int64, string, error)); ok {
		return rf(ctx, info, password)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.UserInfo, string) int64); ok {
		r0 = rf(ctx, info, password)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.UserInfo, string) string); ok {
		r1 = rf(ctx, info, password)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(context.Context, *model.UserInfo, string) error); ok {
		r2 = rf(ctx, info, password)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Get provides a mock function with given fields: ctx, id
func (_m *UserRepository) Get(ctx context.Context, id int64) (*model.User, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*model.User, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *model.User); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
