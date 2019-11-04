package app

import "errors"

var ErrRouteNotFound = errors.New(`route not found`)

type Point struct {
	Lng float64 `json:"lng"`
	Lat float64 `json:"lat"`
}

type Car struct {
	ID          int64 `json:"id"`
	Coordinates Point `json:"coordinates"`
}

type Route struct {
	Point
	Time int64 `json:"time"`
}
