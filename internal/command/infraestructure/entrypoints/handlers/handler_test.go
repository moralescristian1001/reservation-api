package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"

	"reservation-api/internal/command/core/domain"
	businessError "reservation-api/internal/command/core/error"
	"reservation-api/internal/command/infraestructure/entrypoints/handlers/contracts"
	"reservation-api/internal/command/infraestructure/entrypoints/handlers/mapper"
	mocks "reservation-api/mocks/command"
)

type deps struct {
	useCase             *mocks.MockUseCase
	requestMapper       mapper.RequestMapper
	apiResponseProvider ApiResponseProvider
}

func makeDependencies(t *testing.T) deps {
	return deps{
		useCase:             mocks.NewMockUseCase(gomock.NewController(t)),
		requestMapper:       mapper.RequestMapper{},
		apiResponseProvider: ApiResponseProvider{apiError: APIError{}},
	}
}

func TestGetReservations(t *testing.T) {
	reservation := domain.NewReservation("Cliente1", 1)
	type want struct {
		status   int
		response any
	}
	tests := []struct {
		name      string
		body      contracts.ReservationRequest
		mocks     func(deps)
		want      func() *want
		wantError bool
	}{
		{
			name: "Get reservations successfully",
			body: contracts.ReservationRequest{ClientID: "Cliente1", Reservation: 1},
			mocks: func(d deps) {
				d.useCase.EXPECT().SaveReservation(*reservation).Return(nil)
			},
			want: func() *want {
				return &want{status: http.StatusCreated, response: `{"message":"The reservation has been successfully created."}`}
			},
			wantError: false,
		},
		{
			name: "Fails trying to get reservations when client already exist",
			body: contracts.ReservationRequest{ClientID: "Cliente1", Reservation: 1},
			mocks: func(d deps) {
				err := businessError.NewClientAlreadyHasReservation("error saving the reservation")
				d.useCase.EXPECT().SaveReservation(*reservation).Return(err)
			},
			want: func() *want {
				return &want{status: http.StatusConflict, response: `{"status":409,"message":"error saving the reservation"}`}
			},
			wantError: true,
		},
		{
			name: "Fails trying to get reservations",
			body: contracts.ReservationRequest{ClientID: "Cliente1", Reservation: 1},
			mocks: func(d deps) {
				d.useCase.EXPECT().SaveReservation(*reservation).Return(errors.New("error"))
			},
			want: func() *want {
				return &want{status: http.StatusInternalServerError, response: `{"status":500,"message":"error"}`}
			},
			wantError: true,
		},
		{
			name:  "Error obtaining request data",
			body:  contracts.ReservationRequest{},
			mocks: func(d deps) {},
			want: func() *want {
				return &want{status: http.StatusBadRequest, response: `{"error":"Key: 'ReservationRequest.ClientID' Error:Field validation for 'ClientID' failed on the 'required' tag\nKey: 'ReservationRequest.Reservation' Error:Field validation for 'Reservation' failed on the 'required' tag"}`}
			},
			wantError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dep := makeDependencies(t)
			tt.mocks(dep)
			want := tt.want()
			handler := NewReservationHandler(dep.useCase, dep.requestMapper, dep.apiResponseProvider)
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			jsonBytes, err := json.Marshal(tt.body)
			if err != nil {
				panic(err)
			}
			ctx.Request = &http.Request{
				Header: make(http.Header),
				URL:    &url.URL{},
			}
			ctx.Request.Method = "POST"
			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(jsonBytes))
			handler.HandleReservation(ctx)

			body := strings.TrimSuffix(w.Body.String(), "\n")
			assert.Equal(t, want.response, body)
			assert.Equal(t, want.status, w.Code)

		})
	}
}
