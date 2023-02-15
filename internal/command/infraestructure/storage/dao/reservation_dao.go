package dao

import (
	"gorm.io/gorm"
	"reservation-api/internal/command/core/domain"
	businessError "reservation-api/internal/command/core/error"
	"reservation-api/internal/command/infraestructure/storage/dto"
	"strings"
)

type PostgresReservationDAO struct {
	gormDB *gorm.DB
	reservationMapper ReservationMapper
}

type ReservationMapper interface {
	ToDTO(reservation domain.Reservation) dto.Reservation
}

func NewPostgresReservationDAO(gormDB *gorm.DB, reservationMapper ReservationMapper) *PostgresReservationDAO {
	return &PostgresReservationDAO{gormDB: gormDB, reservationMapper: reservationMapper}
}

func (p PostgresReservationDAO) Save(reservation domain.Reservation) error {

	result := p.gormDB.Create(p.reservationMapper.ToDTO(reservation))

	if result.Error != nil {
		return analyzeError(result.Error.Error())
	}

	return nil

}

func analyzeError(error string) error {
	if strings.Contains(error, "idx_reservations_client_id") {
		return businessError.NewClientAlreadyHasReservation("The client already has a reservation.")
	}else if strings.Contains(error, "idx_reservations_reservation") {
		return businessError.NewNumberAlreadyReserved("Number already reserved.")
	}else{
		return businessError.NewPersistenceError("An error occurred while saving the reservation, try again later.")
	}
}

