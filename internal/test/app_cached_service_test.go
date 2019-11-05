package test

import (
	"context"
	"errors"
	"testing"

	"github.com/asardak/arrival-time-service/internal/test/mocks"

	"github.com/stretchr/testify/assert"

	"github.com/asardak/arrival-time-service/internal/app"
)

func TestCachedService_GetArrivalTime(t *testing.T) {
	t.Run("Cache hit", func(t *testing.T) {
		service := &mocks.ArrivalTimeService{}
		repo := &mocks.Repository{}

		point := app.Point{Lat: 55.112233, Lng: 37.112233}
		route := app.Route{
			Point: point,
			Time:  10,
		}
		ctx := context.Background()

		repo.On("FindRoute", ctx, point).Return(route, nil)

		cachedService := app.NewCachedService(service, repo)
		res, err := cachedService.GetArrivalTime(ctx, point)

		assert.Nil(t, err)
		assert.Equal(t, int64(10), res)
	})

	t.Run("Cache miss", func(t *testing.T) {
		service := &mocks.ArrivalTimeService{}
		repo := &mocks.Repository{}

		point := app.Point{Lat: 55.112233, Lng: 37.112233}
		route := app.Route{
			Point: point,
			Time:  10,
		}
		ctx := context.Background()

		repo.On("FindRoute", ctx, point).Return(app.Route{}, app.ErrRouteNotFound)
		repo.On("SaveRoute", ctx, route).Return(nil)
		service.On("GetArrivalTime", ctx, point).Return(int64(10), nil)

		cachedService := app.NewCachedService(service, repo)
		res, err := cachedService.GetArrivalTime(ctx, point)

		assert.Nil(t, err)
		assert.Equal(t, int64(10), res)
	})

	t.Run("Error", func(t *testing.T) {
		service := &mocks.ArrivalTimeService{}
		repo := &mocks.Repository{}

		point := app.Point{Lat: 55.112233, Lng: 37.112233}
		ctx := context.Background()

		repo.On("FindRoute", ctx, point).Return(app.Route{}, app.ErrRouteNotFound)
		service.On("GetArrivalTime", ctx, point).Return(int64(0), errors.New("error"))

		cachedService := app.NewCachedService(service, repo)
		_, err := cachedService.GetArrivalTime(ctx, point)

		assert.NotNil(t, err)
	})
}
