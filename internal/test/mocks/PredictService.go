// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	app "github.com/asardak/arrival-time-service/internal/app"

	mock "github.com/stretchr/testify/mock"
)

// PredictService is an autogenerated mock type for the PredictService type
type PredictService struct {
	mock.Mock
}

// GetTime provides a mock function with given fields: ctx, point, cars
func (_m *PredictService) GetTime(ctx context.Context, point app.Point, cars []*app.Car) ([]int64, error) {
	ret := _m.Called(ctx, point, cars)

	var r0 []int64
	if rf, ok := ret.Get(0).(func(context.Context, app.Point, []*app.Car) []int64); ok {
		r0 = rf(ctx, point, cars)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int64)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, app.Point, []*app.Car) error); ok {
		r1 = rf(ctx, point, cars)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
