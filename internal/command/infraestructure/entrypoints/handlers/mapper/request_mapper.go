package mapper

import (
	"reservation-api/internal/command/core/domain"
	"reservation-api/internal/command/infraestructure/entrypoints/handlers/contracts"
)

type RequestMapper struct {}

func (r RequestMapper) ToDomain(request contracts.ReservationRequest) domain.Reservation  {
	return *domain.NewReservation(request.ClientID, request.Reservation)
}
