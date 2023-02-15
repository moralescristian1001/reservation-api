package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reservation-api/internal/command/core/domain"
	"reservation-api/internal/command/infraestructure/entrypoints/handlers/contracts"
)

type UseCase interface {
	SaveReservation(reservation domain.Reservation) error
}

type RequestMapper interface {
	ToDomain(request contracts.ReservationRequest) domain.Reservation
}

type ReservationHandler struct {
	useCase             UseCase
	requestMapper RequestMapper
	apiResponseProvider ApiResponseProvider
}

func NewReservationHandler(useCase UseCase, requestMapper RequestMapper, apiResponseProvider ApiResponseProvider) *ReservationHandler {
	return &ReservationHandler{useCase: useCase, requestMapper: requestMapper, apiResponseProvider: apiResponseProvider}
}


func (r ReservationHandler) HandleReservation(c *gin.Context) {

	var reservationRequest contracts.ReservationRequest

	if err := c.ShouldBindJSON(&reservationRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := r.useCase.SaveReservation(r.requestMapper.ToDomain(reservationRequest))
	if err != nil {
		_ = r.apiResponseProvider.MapperAPIError(c.Writer, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "The reservation was successfully completed.",
	})

	return
}
