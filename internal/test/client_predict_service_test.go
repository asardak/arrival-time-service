package test

import (
	"context"
	"errors"
	"testing"

	"github.com/asardak/arrival-time-service/internal/test/mocks"

	"github.com/stretchr/testify/assert"

	"github.com/asardak/arrival-time-service/pkg/predict-service/client/operations"
	"github.com/asardak/arrival-time-service/pkg/predict-service/models"

	"github.com/asardak/arrival-time-service/internal/app"

	"github.com/asardak/arrival-time-service/internal/pkg/client"
)

func TestPredictService_GetTimeMin(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		predictService := &mocks.PredictServiceClient{}

		ctx := context.Background()
		cars := []*app.Car{
			{ID: 111, Coordinates: app.Point{Lat: 56.112233, Lng: 38.112233}},
			{ID: 222, Coordinates: app.Point{Lat: 57.112233, Lng: 39.112233}},
			{ID: 333, Coordinates: app.Point{Lat: 58.112233, Lng: 40.112233}},
		}
		point := app.Point{Lat: 55.112233, Lng: 37.112233}
		req := &operations.PredictParams{
			PositionList: operations.PredictBody{
				Source: []models.Position{
					{Lat: 56.112233, Lng: 38.112233},
					{Lat: 57.112233, Lng: 39.112233},
					{Lat: 58.112233, Lng: 40.112233},
				},
				Target: models.Position{
					Lat: point.Lat,
					Lng: point.Lng,
				},
			},
			Context: ctx,
		}
		resp := &operations.PredictOK{
			Payload: []int64{10, 20, 30},
		}

		predictService.On("Predict", req).Return(resp, nil)

		service := client.NewPredictService(predictService)

		res, err := service.GetTime(ctx, point, cars)

		assert.Nil(t, err)
		assert.Equal(t, []int64{10, 20, 30}, res)
	})

	t.Run("Error", func(t *testing.T) {
		predictService := &mocks.PredictServiceClient{}

		ctx := context.Background()
		cars := []*app.Car{
			{ID: 111, Coordinates: app.Point{Lat: 56.112233, Lng: 38.112233}},
			{ID: 222, Coordinates: app.Point{Lat: 57.112233, Lng: 39.112233}},
			{ID: 333, Coordinates: app.Point{Lat: 58.112233, Lng: 40.112233}},
		}
		point := app.Point{Lat: 55.112233, Lng: 37.112233}
		req := &operations.PredictParams{
			PositionList: operations.PredictBody{
				Source: []models.Position{
					{Lat: 56.112233, Lng: 38.112233},
					{Lat: 57.112233, Lng: 39.112233},
					{Lat: 58.112233, Lng: 40.112233},
				},
				Target: models.Position{
					Lat: point.Lat,
					Lng: point.Lng,
				},
			},
			Context: ctx,
		}

		predictService.On("Predict", req).Return(nil, errors.New("error"))

		service := client.NewPredictService(predictService)

		_, err := service.GetTime(ctx, point, cars)

		assert.NotNil(t, err)
	})
}
