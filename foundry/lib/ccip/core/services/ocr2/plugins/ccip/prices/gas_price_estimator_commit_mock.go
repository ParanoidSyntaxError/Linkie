// Code generated by mockery v2.28.1. DO NOT EDIT.

package prices

import (
	context "context"
	big "math/big"

	mock "github.com/stretchr/testify/mock"
)

// MockGasPriceEstimatorCommit is an autogenerated mock type for the GasPriceEstimatorCommit type
type MockGasPriceEstimatorCommit struct {
	mock.Mock
}

// DenoteInUSD provides a mock function with given fields: p, wrappedNativePrice
func (_m *MockGasPriceEstimatorCommit) DenoteInUSD(p GasPrice, wrappedNativePrice *big.Int) (GasPrice, error) {
	ret := _m.Called(p, wrappedNativePrice)

	var r0 GasPrice
	var r1 error
	if rf, ok := ret.Get(0).(func(GasPrice, *big.Int) (GasPrice, error)); ok {
		return rf(p, wrappedNativePrice)
	}
	if rf, ok := ret.Get(0).(func(GasPrice, *big.Int) GasPrice); ok {
		r0 = rf(p, wrappedNativePrice)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(GasPrice)
		}
	}

	if rf, ok := ret.Get(1).(func(GasPrice, *big.Int) error); ok {
		r1 = rf(p, wrappedNativePrice)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Deviates provides a mock function with given fields: p1, p2
func (_m *MockGasPriceEstimatorCommit) Deviates(p1 GasPrice, p2 GasPrice) (bool, error) {
	ret := _m.Called(p1, p2)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(GasPrice, GasPrice) (bool, error)); ok {
		return rf(p1, p2)
	}
	if rf, ok := ret.Get(0).(func(GasPrice, GasPrice) bool); ok {
		r0 = rf(p1, p2)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(GasPrice, GasPrice) error); ok {
		r1 = rf(p1, p2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGasPrice provides a mock function with given fields: ctx
func (_m *MockGasPriceEstimatorCommit) GetGasPrice(ctx context.Context) (GasPrice, error) {
	ret := _m.Called(ctx)

	var r0 GasPrice
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (GasPrice, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) GasPrice); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(GasPrice)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Median provides a mock function with given fields: gasPrices
func (_m *MockGasPriceEstimatorCommit) Median(gasPrices []GasPrice) (GasPrice, error) {
	ret := _m.Called(gasPrices)

	var r0 GasPrice
	var r1 error
	if rf, ok := ret.Get(0).(func([]GasPrice) (GasPrice, error)); ok {
		return rf(gasPrices)
	}
	if rf, ok := ret.Get(0).(func([]GasPrice) GasPrice); ok {
		r0 = rf(gasPrices)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(GasPrice)
		}
	}

	if rf, ok := ret.Get(1).(func([]GasPrice) error); ok {
		r1 = rf(gasPrices)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// String provides a mock function with given fields: p
func (_m *MockGasPriceEstimatorCommit) String(p GasPrice) string {
	ret := _m.Called(p)

	var r0 string
	if rf, ok := ret.Get(0).(func(GasPrice) string); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTNewMockGasPriceEstimatorCommit interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockGasPriceEstimatorCommit creates a new instance of MockGasPriceEstimatorCommit. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockGasPriceEstimatorCommit(t mockConstructorTestingTNewMockGasPriceEstimatorCommit) *MockGasPriceEstimatorCommit {
	mock := &MockGasPriceEstimatorCommit{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
