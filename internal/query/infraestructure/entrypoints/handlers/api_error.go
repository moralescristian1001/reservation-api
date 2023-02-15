package handlers

type APIError struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
}

func (a *APIError) Error() string {
	return a.Message
}

func (a *APIError) setAPIError(err error, status int) {
	a.Status = status
	a.Message = err.Error()
}
