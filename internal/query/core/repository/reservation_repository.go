package repository

import (
	"reservation-api/internal/query/infraestructure/storage/dto"
)

type ReservationRepository struct {
	reservationDAO ReservationDAO
}

type ReservationDAO interface {
	GetAllReservations() ([]dto.Reservation, error)
}

func NewReservationRepository(reservationDAO ReservationDAO) *ReservationRepository {
	return &ReservationRepository{reservationDAO: reservationDAO}
}

func (r ReservationRepository) GetAllReservations() ([]dto.Reservation, error) {

	return r.reservationDAO.GetAllReservations()
}


