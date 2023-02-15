package error

type PersistenceError struct {
	message string
}

func NewPersistenceError(message string) PersistenceError {
	return PersistenceError{message: message}
}

func (p PersistenceError) Error() string {
	return p.message
}
