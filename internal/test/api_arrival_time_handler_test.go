package test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/mock"

	"github.com/asardak/arrival-time-service/internal/app"

	"github.com/asardak/arrival-time-service/internal/pkg/api"
	"github.com/asardak/arrival-time-service/internal/test/mocks"
)

func TestArrivalTimeHandler_Endpoint(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		ctx := context.Background()
		eCtx := &mocks.EchoContextMock{}
		service := &mocks.ArrivalTimeService{}

		httpReq := &http.Request{}
		httpReq = httpReq.WithContext(ctx)
		req := &api.ArrivalTimeRequest{Lat: 55.112233, Lng: 37.112233}
		point := app.Point{Lat: 55.112233, Lng: 37.112233}
		result := int64(123)
		resp := api.ArrivalTimeResponse{ArrivalTimeMinutes: result}

		eCtx.On("Bind", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			arg := args.Get(0).(*api.ArrivalTimeRequest)
			*arg = *req
		})
		eCtx.On("Request").Return(httpReq)
		service.On("GetArrivalTime", ctx, point).Return(result, nil)
		eCtx.On("JSON", mock.Anything, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			assert.Equal(t, http.StatusOK, args.Get(0).(int))
			assert.Equal(t, resp, args.Get(1).(api.ArrivalTimeResponse))
		})

		h := api.NewArrivalTimeHandler(service)
		err := h.Endpoint(eCtx)

		assert.Nil(t, err)
	})

	t.Run("Validation", func(t *testing.T) {
		ctx := context.Background()
		eCtx := &mocks.EchoContextMock{}
		service := &mocks.ArrivalTimeService{}

		httpReq := &http.Request{}
		httpReq = httpReq.WithContext(ctx)
		req := &api.ArrivalTimeRequest{Lat: 91, Lng: 181}

		eCtx.On("Bind", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			arg := args.Get(0).(*api.ArrivalTimeRequest)
			*arg = *req
		})
		eCtx.On("Request").Return(httpReq)

		h := api.NewArrivalTimeHandler(service)
		err := h.Endpoint(eCtx)

		assert.NotNil(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		ctx := context.Background()
		eCtx := &mocks.EchoContextMock{}
		service := &mocks.ArrivalTimeService{}

		httpReq := &http.Request{}
		httpReq = httpReq.WithContext(ctx)
		req := &api.ArrivalTimeRequest{Lat: 55.112233, Lng: 37.112233}
		point := app.Point{Lat: 55.112233, Lng: 37.112233}

		eCtx.On("Bind", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			arg := args.Get(0).(*api.ArrivalTimeRequest)
			*arg = *req
		})
		eCtx.On("Request").Return(httpReq)
		service.On("GetArrivalTime", ctx, point).Return(int64(0), errors.New("error"))

		h := api.NewArrivalTimeHandler(service)
		err := h.Endpoint(eCtx)

		assert.NotNil(t, err)
	})
}
