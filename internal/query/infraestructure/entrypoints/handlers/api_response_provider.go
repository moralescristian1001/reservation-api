package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	businessError "reservation-api/internal/query/core/error"
)

type ApiResponseProvider struct {
	apiError APIError
}

func (r *ApiResponseProvider) MapperAPIError(w http.ResponseWriter, err error) error {
	switch err.(type) {
	case businessError.ReadError:
		r.apiError.setAPIError(err, http.StatusInternalServerError)
	default:
		r.apiError.setAPIError(err, http.StatusInternalServerError)
	}

	w.Header().Add("statusCode", fmt.Sprint(r.apiError.Status))
	_ = r.encodeJSON(w, r.apiError, r.apiError.Status)
	return nil
}

func (r ApiResponseProvider) encodeJSON(w http.ResponseWriter, v interface{}, code int) error {
	for k, values := range w.Header() {
		for _, v := range values {
			w.Header().Add(k, v)
		}
	}

	var jsonData []byte

	var err error
	switch v := v.(type) {
	case []byte:
		jsonData = v
	case io.Reader:
		jsonData, err = io.ReadAll(v)
	default:
		jsonData, err = json.Marshal(v)
	}

	if err != nil {
		return err
	}

	// Set the content type.
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// Write the status code to the response and context.
	w.WriteHeader(code)

	// Send the result back to the client.
	if _, err := w.Write(jsonData); err != nil {
		return err
	}

	return nil
}


