// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Car car
// swagger:model car
type Car struct {

	// Car id
	// Required: true
	// Minimum: 1
	ID int64 `json:"id"`

	// Car current latitude
	// Required: true
	// Maximum: 90
	// Minimum: -90
	Lat float64 `json:"lat"`

	// Car current longitude
	// Required: true
	// Maximum: 180
	// Minimum: -180
	Lng float64 `json:"lng"`
}

// Validate validates this car
func (m *Car) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLng(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Car) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", int64(m.ID)); err != nil {
		return err
	}

	if err := validate.Minimum("id", "body", float64(m.ID), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *Car) validateLat(formats strfmt.Registry) error {

	if err := validate.Required("lat", "body", float64(m.Lat)); err != nil {
		return err
	}

	if err := validate.Minimum("lat", "body", float64(m.Lat), -90, false); err != nil {
		return err
	}

	if err := validate.Maximum("lat", "body", float64(m.Lat), 90, false); err != nil {
		return err
	}

	return nil
}

func (m *Car) validateLng(formats strfmt.Registry) error {

	if err := validate.Required("lng", "body", float64(m.Lng)); err != nil {
		return err
	}

	if err := validate.Minimum("lng", "body", float64(m.Lng), -180, false); err != nil {
		return err
	}

	if err := validate.Maximum("lng", "body", float64(m.Lng), 180, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Car) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Car) UnmarshalBinary(b []byte) error {
	var res Car
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
