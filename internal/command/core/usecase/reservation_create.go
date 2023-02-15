package usecase

import (
	"reservation-api/internal/command/core/domain"
)

type CreateUseCase struct {
	reservationRepository ReservationRepository
}

type ReservationRepository interface {
	Save(reservation domain.Reservation) error
}

func NewCreateUseCase(reservationRepository ReservationRepository) *CreateUseCase {
	return &CreateUseCase{reservationRepository: reservationRepository}
}

func (c CreateUseCase) SaveReservation(reservation domain.Reservation) error {

	err := c.reservationRepository.Save(reservation)
	if err != nil {
		return err
	}

	return nil
}
