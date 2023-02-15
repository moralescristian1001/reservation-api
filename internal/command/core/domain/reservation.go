package domain

type Reservation struct {
	clientID string
	reservation int
}

func NewReservation(clientID string, reservation int) *Reservation {
	return &Reservation{clientID: clientID, reservation: reservation}
}

func (r Reservation) ClientID() string {
	return r.clientID
}

func (r Reservation) Reservation() int {
	return r.reservation
}


