package dto

type Reservation struct {
	ClientID string `json:"client_id" gorm:"uniqueIndex:idx_reservations_client_id"`
	Reservation   int    `json:"reservation" gorm:"uniqueIndex:idx_reservations_reservation"`
}

func NewReservation(clientID string, reservation int) *Reservation {
	return &Reservation{ClientID: clientID, Reservation: reservation}
}