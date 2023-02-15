package error

type ReadError struct {
	message string
}

func NewReadError(message string) ReadError {
	return ReadError{message: message}
}

func (r ReadError) Error() string {
	return r.message
}
