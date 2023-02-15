package repository

import "reservation-api/internal/command/core/domain"

type ReservationRepository struct {
	reservationDAO ReservationDAO
}

type ReservationDAO interface {
	Save(reservation domain.Reservation) error
}

func NewReservationRepository(reservationDAO ReservationDAO) *ReservationRepository {
	return &ReservationRepository{reservationDAO: reservationDAO}
}

func (r ReservationRepository) Save(reservation domain.Reservation) error {

	return r.reservationDAO.Save(reservation)
}


