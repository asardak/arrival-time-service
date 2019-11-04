package api

import "github.com/labstack/echo/v4"

type Router struct {
	service ArrivalTimeService
}

func NewRouter(service ArrivalTimeService) *Router {
	return &Router{
		service: service,
	}
}

func (r *Router) Mount(e *echo.Echo) {
	e.POST("/arrival-time", NewArrivalTimeHandler(r.service).Endpoint)
}
