package domain

import (
	"errors"
	"net/http"
)

type Err struct {
	error      error
	StatusCode int
	Parameters map[string]string
}

func (e Err) Error() string {
	return e.error.Error()
}

func NewError(message string) *Err {
	return &Err{error: errors.New(message), StatusCode: http.StatusUnprocessableEntity}
}
