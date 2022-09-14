package errs

import "net/http"

type AppErr struct {
	Code    int `json:",omitempty"`
	Message string
}

func (e AppErr) AsMessage() *AppErr {
	return &AppErr{Message: e.Message}
}

func NewNotFoundError(message string) *AppErr {
	return &AppErr{http.StatusNotFound, message}
}

func NewUnexpectedError(message string) *AppErr {
	return &AppErr{http.StatusInternalServerError, message}
}
