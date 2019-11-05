package test

import (
	"context"
	"errors"
	"testing"

	"github.com/asardak/arrival-time-service/internal/test/mocks"

	"github.com/asardak/arrival-time-service/internal/app"

	"github.com/stretchr/testify/assert"
)

func TestArrivalTimeService_GetArrivalTime(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		carService := &mocks.CarService{}
		predictService := &mocks.PredictService{}

		point := app.Point{Lat: 55.112233, Lng: 37.112233}
		cars := []*app.Car{
			{ID: 111, Coordinates: app.Point{Lat: 56.112233, Lng: 38.112233}},
			{ID: 333, Coordinates: app.Point{Lat: 57.112233, Lng: 39.112233}},
			{ID: 333, Coordinates: app.Point{Lat: 58.112233, Lng: 40.112233}},
		}
		ctx := context.Background()
		times := []int64{10, 20, 30}

		carService.On("GetNearestCars", ctx, point).Return(cars, nil)
		predictService.On("GetTime", ctx, point, cars).Return(times, nil)

		service := app.NewArrivalTimeService(carService, predictService)
		res, err := service.GetArrivalTime(ctx, point)

		assert.Nil(t, err)
		assert.Equal(t, int64(10), res)
	})

	t.Run("Error", func(t *testing.T) {
		carService := &mocks.CarService{}
		predictService := &mocks.PredictService{}

		point := app.Point{Lat: 55.112233, Lng: 37.112233}
		cars := []*app.Car{
			{ID: 111, Coordinates: app.Point{Lat: 56.112233, Lng: 38.112233}},
			{ID: 333, Coordinates: app.Point{Lat: 57.112233, Lng: 39.112233}},
			{ID: 333, Coordinates: app.Point{Lat: 58.112233, Lng: 40.112233}},
		}
		ctx := context.Background()

		carService.On("GetNearestCars", ctx, point).Return(cars, nil)
		predictService.On("GetTime", ctx, point, cars).Return(nil, errors.New("error"))

		service := app.NewArrivalTimeService(carService, predictService)
		_, err := service.GetArrivalTime(ctx, point)

		assert.NotNil(t, err)
	})
}
