package app

import (
	"context"
	"fmt"
	"log"
)

type CachedService struct {
	service ArrivalTimeGetter
	repo    Repository
}

func NewCachedService(service ArrivalTimeGetter, repo Repository) *CachedService {
	return &CachedService{
		service: service,
		repo:    repo,
	}
}

func (s *CachedService) GetArrivalTime(ctx context.Context, point Point) (int64, error) {
	r, err := s.repo.FindRoute(ctx, point)
	if err == nil {
		return r.Time, nil
	}

	if err != ErrRouteNotFound {
		log.Printf(`failed to find route: %v`, err)
	}

	t, err := s.service.GetArrivalTime(ctx, point)
	if err != nil {
		return 0, fmt.Errorf(`failed to get arrival time: %v`, err)
	}

	err = s.repo.SaveRoute(ctx, Route{Point: point, Time: t})
	if err != nil {
		log.Printf(`failed to save route: %v`, err)
	}

	return t, nil
}
