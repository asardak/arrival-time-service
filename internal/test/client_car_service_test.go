package test

import (
	"context"
	"errors"
	"testing"

	"github.com/asardak/arrival-time-service/internal/test/mocks"

	"github.com/stretchr/testify/assert"

	"github.com/asardak/arrival-time-service/pkg/car-service/models"

	"github.com/asardak/arrival-time-service/pkg/car-service/client/operations"

	"github.com/asardak/arrival-time-service/internal/app"

	"github.com/asardak/arrival-time-service/internal/pkg/client"
)

func TestCarService_GetNearestCars(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		carsClient := &mocks.CarsServiceClient{}

		cars := []*app.Car{
			{ID: 111, Coordinates: app.Point{Lat: 56.112233, Lng: 38.112233}},
			{ID: 222, Coordinates: app.Point{Lat: 57.112233, Lng: 39.112233}},
			{ID: 333, Coordinates: app.Point{Lat: 58.112233, Lng: 40.112233}},
		}
		point := app.Point{Lat: 55.112233, Lng: 37.112233}
		ctx := context.Background()
		limit := int64(10)
		req := &operations.GetCarsParams{
			Lat:     point.Lat,
			Lng:     point.Lng,
			Limit:   limit,
			Context: ctx,
		}
		resp := &operations.GetCarsOK{
			Payload: []models.Car{
				{ID: 111, Lat: 56.112233, Lng: 38.112233},
				{ID: 222, Lat: 57.112233, Lng: 39.112233},
				{ID: 333, Lat: 58.112233, Lng: 40.112233},
			},
		}

		carsClient.On("GetCars", req).Return(resp, nil)

		service := client.NewCarService(carsClient, limit)

		result, err := service.GetNearestCars(ctx, point)

		assert.Nil(t, err)
		assert.Equal(t, cars, result)
	})

	t.Run("Error", func(t *testing.T) {
		carsClient := &mocks.CarsServiceClient{}

		point := app.Point{Lat: 55.112233, Lng: 37.112233}
		ctx := context.Background()
		limit := int64(10)
		req := &operations.GetCarsParams{
			Lat:     point.Lat,
			Lng:     point.Lng,
			Limit:   limit,
			Context: ctx,
		}

		carsClient.On("GetCars", req).Return(nil, errors.New("error"))

		service := client.NewCarService(carsClient, limit)

		_, err := service.GetNearestCars(ctx, point)

		assert.NotNil(t, err)
	})
}
