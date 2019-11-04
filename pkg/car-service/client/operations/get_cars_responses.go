// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/asardak/arrival-time-service/pkg/car-service/models"
)

// GetCarsReader is a Reader for the GetCars structure.
type GetCarsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCarsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCarsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCarsOK creates a GetCarsOK with default headers values
func NewGetCarsOK() *GetCarsOK {
	return &GetCarsOK{}
}

/*GetCarsOK handles this case with default header values.

Car list
*/
type GetCarsOK struct {
	Payload []models.Car
}

func (o *GetCarsOK) Error() string {
	return fmt.Sprintf("[GET /cars][%d] getCarsOK  %+v", 200, o.Payload)
}

func (o *GetCarsOK) GetPayload() []models.Car {
	return o.Payload
}

func (o *GetCarsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
