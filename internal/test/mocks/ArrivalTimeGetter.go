// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	app "github.com/asardak/arrival-time-service/internal/app"

	mock "github.com/stretchr/testify/mock"
)

// ArrivalTimeGetter is an autogenerated mock type for the ArrivalTimeGetter type
type ArrivalTimeGetter struct {
	mock.Mock
}

// GetArrivalTime provides a mock function with given fields: ctx, point
func (_m *ArrivalTimeGetter) GetArrivalTime(ctx context.Context, point app.Point) (int64, error) {
	ret := _m.Called(ctx, point)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, app.Point) int64); ok {
		r0 = rf(ctx, point)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, app.Point) error); ok {
		r1 = rf(ctx, point)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
