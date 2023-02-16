package dao

import (
	"gorm.io/gorm"

	businessError "reservation-api/internal/query/core/error"
	"reservation-api/internal/query/infraestructure/storage/dto"
)

type PostgresReservationDAO struct {
	gormDB *gorm.DB
}

func NewPostgresReservationDAO(gormDB *gorm.DB) *PostgresReservationDAO {
	return &PostgresReservationDAO{gormDB: gormDB}
}

func (p PostgresReservationDAO) GetAllReservations() ([]dto.Reservation, error) {

	var reservationsDTO []dto.Reservation
	result := p.gormDB.Find(&reservationsDTO)

	if result.Error != nil {
		return nil, businessError.NewReadError("Error consulting the reservations.")
	}

	return reservationsDTO, nil
}
