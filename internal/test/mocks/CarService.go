// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	app "github.com/asardak/arrival-time-service/internal/app"

	mock "github.com/stretchr/testify/mock"
)

// CarService is an autogenerated mock type for the CarService type
type CarService struct {
	mock.Mock
}

// GetNearestCars provides a mock function with given fields: ctx, point
func (_m *CarService) GetNearestCars(ctx context.Context, point app.Point) ([]*app.Car, error) {
	ret := _m.Called(ctx, point)

	var r0 []*app.Car
	if rf, ok := ret.Get(0).(func(context.Context, app.Point) []*app.Car); ok {
		r0 = rf(ctx, point)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*app.Car)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, app.Point) error); ok {
		r1 = rf(ctx, point)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
