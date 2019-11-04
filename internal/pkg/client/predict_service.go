package client

import (
	"context"
	"fmt"

	"github.com/asardak/arrival-time-service/pkg/predict-service/models"

	"github.com/asardak/arrival-time-service/internal/app"
	"github.com/asardak/arrival-time-service/pkg/predict-service/client/operations"
)

type PredictServiceClient interface {
	Predict(params *operations.PredictParams) (*operations.PredictOK, error)
}

type PredictService struct {
	client PredictServiceClient
}

func NewPredictService(client PredictServiceClient) *PredictService {
	return &PredictService{
		client: client,
	}
}

func (s *PredictService) GetTimeMin(ctx context.Context, point app.Point, cars []*app.Car) ([]int64, error) {
	positions := make([]models.Position, 0, len(cars))
	for _, c := range cars {
		positions = append(positions, models.Position{
			Lat: c.Coordinates.Lat,
			Lng: c.Coordinates.Lng,
		})
	}

	resp, err := s.client.Predict(&operations.PredictParams{
		PositionList: operations.PredictBody{
			Source: positions,
			Target: models.Position{
				Lat: point.Lat,
				Lng: point.Lng,
			},
		},
		Context: ctx,
	})

	if err != nil {
		return nil, fmt.Errorf(`failed get response from predict-service: %w`, err)
	}

	return resp.Payload, nil
}
