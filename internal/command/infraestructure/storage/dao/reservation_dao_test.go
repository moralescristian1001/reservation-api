package dao

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"reservation-api/internal/command/core/domain"
	"reservation-api/internal/command/infraestructure/storage/mapper"
)

func TestGetAllReservations(t *testing.T) {
	sqlInsert := `INSERT INTO "reservations" ("client_id","reservation")`
	tests := []struct {
		name      string
		args      domain.Reservation
		mocks     func(sqlmock.Sqlmock)
		wantError bool
	}{
		{
			name: "Save a reservation successfully",
			mocks: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(sqlInsert)).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			wantError: false,
		},
		{
			name: "Failure to save a reservation when the client already has a reservation",
			mocks: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(sqlInsert)).WillReturnError(errors.New(`duplicate key value violates unique constraint "idx_reservations_client_id"`))
			},
			wantError: true,
		},
		{
			name: "Failure when trying to save a reservation when the reservation number already exists",
			mocks: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(sqlInsert)).WillReturnError(errors.New(`duplicate key value violates unique constraint "idx_reservations_reservation"`))
			},
			wantError: true,
		},
		{
			name: "Failure when trying to save a reservation",
			mocks: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(sqlInsert)).WillReturnError(errors.New(`an error has occurred`))
			},
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			defer db.Close()

			dialector := postgres.New(postgres.Config{
				DSN:                  "sqlmock_db_0",
				DriverName:           "postgres",
				Conn:                 db,
				PreferSimpleProtocol: true,
			})
			gdb, _ := gorm.Open(dialector, &gorm.Config{})
			reservationMapper := mapper.ReservationMapper{}
			var dao = &PostgresReservationDAO{gormDB: gdb, reservationMapper: reservationMapper}
			tt.mocks(mock)

			err := dao.Save(tt.args)
			if (err != nil) != tt.wantError {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantError)
				return
			}
		})
	}
}
