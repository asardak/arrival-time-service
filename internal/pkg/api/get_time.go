package api

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/asardak/arrival-time-service/internal/app"
)

type ArrivalTimeService interface {
	GetArrivalTime(ctx context.Context, point app.Point) (int64, error)
}

type ArrivalTimeHandler struct {
	service ArrivalTimeService
}

type ArrivalTimeRequest struct {
	Lng float64 `json:"lng"`
	Lat float64 `json:"lat"`
}

func (r *ArrivalTimeRequest) Validate() error {
	if r.Lng < -180 || r.Lng > 180 {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "lng has invalid value")
	}

	if r.Lat < -90 || r.Lat > 90 {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "lat has invalid value")
	}

	return nil
}

type ArrivalTimeResponse struct {
	ArrivalTimeMinutes int64 `json:"arrival_time_minutes"`
}

func NewArrivalTimeHandler(service ArrivalTimeService) *ArrivalTimeHandler {
	return &ArrivalTimeHandler{service: service}
}

func (e *ArrivalTimeHandler) Endpoint(c echo.Context) error {
	ctx := c.Request().Context()

	var req ArrivalTimeRequest
	err := c.Bind(&req)
	if err != nil {
		return err
	}

	err = req.Validate()
	if err != nil {
		return err
	}

	point := app.Point{Lng: req.Lng, Lat: req.Lat}
	result, err := e.service.GetArrivalTime(ctx, point)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, ArrivalTimeResponse{result})
}
