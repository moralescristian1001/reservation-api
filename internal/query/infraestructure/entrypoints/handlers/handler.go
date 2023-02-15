package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reservation-api/internal/query/infraestructure/storage/dto"
)

type ReservationRepository interface {
	GetAllReservations() ([]dto.Reservation, error)
}

type ReservationHandler struct {
	reservationRepository ReservationRepository
	apiResponseProvider ApiResponseProvider
}

func NewReservationHandler(reservationRepository ReservationRepository, apiResponseProvider ApiResponseProvider) *ReservationHandler {
	return &ReservationHandler{reservationRepository: reservationRepository, apiResponseProvider: apiResponseProvider}
}

func (r ReservationHandler) GetReservations(c *gin.Context) {

	reservations, err := r.reservationRepository.GetAllReservations()
	if err != nil {
		_ = r.apiResponseProvider.MapperAPIError(c.Writer, err)
		return
	}

	c.JSON(http.StatusOK, reservations)

	return
}
