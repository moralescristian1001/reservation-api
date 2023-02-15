package error

type ClientAlreadyHasReservation struct {
	message string
}

func NewClientAlreadyHasReservation(message string) ClientAlreadyHasReservation {
	return ClientAlreadyHasReservation{message: message}
}

func (c ClientAlreadyHasReservation) Error() string {
	return c.message
}
