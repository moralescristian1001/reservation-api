package dao

import (
	"errors"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"reservation-api/internal/query/infraestructure/storage/dto"
)

func TestGetAllReservations(t *testing.T) {
	testQuery := `SELECT * FROM "reservations"`
	tests := []struct {
		name      string
		wantError bool
		mocks     func(sqlmock.Sqlmock)
		want      []dto.Reservation
	}{
		{
			name: "Get all reservations successfully",
			mocks: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"client_id", "reservation"}).
					AddRow("Cliente1", 1).AddRow("Cliente2", 10).AddRow("Cliente3", 22)
				mock.ExpectQuery(regexp.QuoteMeta(testQuery)).WillReturnRows(rows)
			},
			want:      []dto.Reservation{{ClientID: "Cliente1", Reservation: 1}, {ClientID: "Cliente2", Reservation: 10}, {ClientID: "Cliente3", Reservation: 22}},
			wantError: false,
		},
		{
			name: "Fails trying to get all reservations",
			mocks: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(testQuery)).WillReturnError(errors.New("error"))
			},
			want:      nil,
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
			var dao = &PostgresReservationDAO{gormDB: gdb}
			tt.mocks(mock)

			got, err := dao.GetAllReservations()
			if (err != nil) != tt.wantError {
				t.Errorf("GetAllReservations() error = %v, wantErr %v", err, tt.wantError)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expected = %v, result %v", tt.want, got)
			}
		})
	}
}
