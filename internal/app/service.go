package app

import (
	"context"
	"errors"
	"fmt"
)

type ArrivalTimeService struct {
	cars       CarService
	prediction PredictService
}

func NewArrivalTimeService(cars CarService, routes PredictService) *ArrivalTimeService {
	return &ArrivalTimeService{
		cars:       cars,
		prediction: routes,
	}
}

func (a *ArrivalTimeService) GetArrivalTime(ctx context.Context, point Point) (int64, error) {
	cars, err := a.cars.GetNearestCars(ctx, point)
	if err != nil {
		return 0, fmt.Errorf(`failed to get nearest cars, err: %w`, err)
	}

	if len(cars) == 0 {
		return 0, errors.New("no cars available")
	}

	durations, err := a.prediction.GetTime(ctx, point, cars)
	if err != nil {
		return 0, fmt.Errorf(`failed to get prediction, err: %w`, err)
	}

	min := durations[0]
	for i := range durations[1:] {
		if durations[i+1] < min {
			min = durations[i+1]
		}
	}

	return min, nil
}
