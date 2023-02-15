package contracts

type ReservationRequest struct {
	ClientID string `json:"client_id" binding:"required"`
	Reservation int `json:"reservation" binding:"required,gte=1"`
}