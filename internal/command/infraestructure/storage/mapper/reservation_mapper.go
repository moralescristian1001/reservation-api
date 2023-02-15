package mapper

import (
	"reservation-api/internal/command/core/domain"
	"reservation-api/internal/command/infraestructure/storage/dto"
)

type ReservationMapper struct {}

func (rm ReservationMapper) ToDTO(reservation domain.Reservation) dto.Reservation {
	return *dto.NewReservation(reservation.ClientID(), reservation.Reservation())
}


