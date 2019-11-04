package app

import (
	"context"
)

type ArrivalTimeGetter interface {
	GetArrivalTime(ctx context.Context, point Point) (int64, error)
}

type Repository interface {
	FindRoute(ctx context.Context, point Point) (Route, error)
	SaveRoute(ctx context.Context, route Route) error
}

type CarService interface {
	GetNearestCars(ctx context.Context, point Point) ([]*Car, error)
}

type PredictService interface {
	GetTimeMin(ctx context.Context, point Point, cars []*Car) ([]int64, error)
}
