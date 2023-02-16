package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"

	"reservation-api/internal/query/infraestructure/storage/dto"
	mocks "reservation-api/mocks/query"
)

type deps struct {
	reservationRepository *mocks.MockReservationRepository
	apiResponseProvider   ApiResponseProvider
}

func makeDependencies(t *testing.T) deps {
	return deps{
		reservationRepository: mocks.NewMockReservationRepository(gomock.NewController(t)),
		apiResponseProvider:   ApiResponseProvider{apiError: APIError{}},
	}
}

func TestGetReservations(t *testing.T) {
	reservations := []dto.Reservation{{ClientID: "Cliente1", Reservation: 1}, {ClientID: "Cliente2", Reservation: 10}, {ClientID: "Cliente3", Reservation: 22}}
	type want struct {
		status   int
		response any
	}
	tests := []struct {
		name      string
		mocks     func(deps)
		want      func() *want
		wantError bool
	}{
		{
			name: "Get reservations successfully",
			mocks: func(d deps) {
				d.reservationRepository.EXPECT().GetAllReservations().Return(reservations, nil)
			},
			want: func() *want {
				return &want{status: http.StatusOK, response: reservations}
			},
			wantError: false,
		},
		{
			name: "Fails trying to get reservations",
			mocks: func(d deps) {
				d.reservationRepository.EXPECT().GetAllReservations().Return(nil, errors.New("error consulting the reservations"))
			},
			want: func() *want {
				return &want{status: http.StatusInternalServerError, response: `{"status":500,"message":"error consulting the reservations"}`}
			},
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dep := makeDependencies(t)
			tt.mocks(dep)
			want := tt.want()
			handler := NewReservationHandler(dep.reservationRepository, dep.apiResponseProvider)
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			handler.GetReservations(ctx)

			assert.Equal(t, want.status, w.Code)
			if tt.wantError {
				body := strings.TrimSuffix(w.Body.String(), "\n")
				assert.Equal(t, want.response, body)
			} else {
				var gotReservations []dto.Reservation
				json.NewDecoder(w.Body).Decode(&gotReservations)
				if !reflect.DeepEqual(gotReservations, want.response) {
					t.Errorf("expected = %v, result %v", want.response, gotReservations)
				}
			}
		})
	}
}
