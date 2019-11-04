package client

import (
	"context"
	"fmt"

	"github.com/asardak/arrival-time-service/internal/app"

	"github.com/asardak/arrival-time-service/pkg/car-service/client/operations"
)

type CarsServiceClient interface {
	GetCars(params *operations.GetCarsParams) (*operations.GetCarsOK, error)
}

type CarService struct {
	client CarsServiceClient
	limit  int64
}

func NewCarService(client CarsServiceClient, limit int64) *CarService {
	return &CarService{
		client: client,
		limit:  limit,
	}
}

func (s *CarService) GetNearestCars(ctx context.Context, point app.Point) ([]*app.Car, error) {
	resp, err := s.client.GetCars(&operations.GetCarsParams{
		Lat:     point.Lat,
		Lng:     point.Lng,
		Limit:   s.limit,
		Context: ctx,
	})

	if err != nil {
		return nil, fmt.Errorf(`failed to get cars from service: %w`, err)
	}

	cars := make([]*app.Car, 0, len(resp.Payload))
	for i := range resp.Payload {
		cars = append(cars, &app.Car{
			ID: resp.Payload[i].ID,
			Coordinates: app.Point{
				Lat: resp.Payload[i].Lat,
				Lng: resp.Payload[i].Lng,
			},
		})
	}

	return cars, nil
}
