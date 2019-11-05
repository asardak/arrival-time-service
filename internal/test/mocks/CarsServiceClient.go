// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	operations "github.com/asardak/arrival-time-service/pkg/car-service/client/operations"
	mock "github.com/stretchr/testify/mock"
)

// CarsServiceClient is an autogenerated mock type for the CarsServiceClient type
type CarsServiceClient struct {
	mock.Mock
}

// GetCars provides a mock function with given fields: params
func (_m *CarsServiceClient) GetCars(params *operations.GetCarsParams) (*operations.GetCarsOK, error) {
	ret := _m.Called(params)

	var r0 *operations.GetCarsOK
	if rf, ok := ret.Get(0).(func(*operations.GetCarsParams) *operations.GetCarsOK); ok {
		r0 = rf(params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*operations.GetCarsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*operations.GetCarsParams) error); ok {
		r1 = rf(params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
