package error

type NumberAlreadyReserved struct {
	message string
}

func NewNumberAlreadyReserved(message string) NumberAlreadyReserved {
	return NumberAlreadyReserved{message: message}
}

func (n NumberAlreadyReserved) Error() string {
	return n.message
}
