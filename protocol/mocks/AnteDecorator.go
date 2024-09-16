// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	types "github.com/cosmos/cosmos-sdk/types"
	mock "github.com/stretchr/testify/mock"
)

// AnteDecorator is an autogenerated mock type for the AnteDecorator type
type AnteDecorator struct {
	mock.Mock
}

// AnteHandle provides a mock function with given fields: ctx, tx, simulate, next
func (_m *AnteDecorator) AnteHandle(ctx types.Context, tx types.Tx, simulate bool, next types.AnteHandler) (types.Context, error) {
	ret := _m.Called(ctx, tx, simulate, next)

	if len(ret) == 0 {
		panic("no return value specified for AnteHandle")
	}

	var r0 types.Context
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, types.Tx, bool, types.AnteHandler) (types.Context, error)); ok {
		return rf(ctx, tx, simulate, next)
	}
	if rf, ok := ret.Get(0).(func(types.Context, types.Tx, bool, types.AnteHandler) types.Context); ok {
		r0 = rf(ctx, tx, simulate, next)
	} else {
		r0 = ret.Get(0).(types.Context)
	}

	if rf, ok := ret.Get(1).(func(types.Context, types.Tx, bool, types.AnteHandler) error); ok {
		r1 = rf(ctx, tx, simulate, next)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAnteDecorator creates a new instance of AnteDecorator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAnteDecorator(t interface {
	mock.TestingT
	Cleanup(func())
}) *AnteDecorator {
	mock := &AnteDecorator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
