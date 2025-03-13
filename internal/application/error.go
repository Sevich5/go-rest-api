package application

import (
	"app/internal/domain"
	"errors"
	"net/http"
)

type Err struct {
	domain.Err
}

func NewError(message string) *Err {
	err := &Err{Err: *domain.NewError(message)}
	err.StatusCode = http.StatusInternalServerError
	return err
}

func NewErrorFromErr(err error) *Err {
	var domainErr *domain.Err
	switch {
	case errors.As(err, &domainErr):
		appErr := NewError(err.Error())
		appErr.StatusCode = domainErr.StatusCode
		appErr.Parameters = domainErr.Parameters
		return appErr
	}
	return NewError(err.Error())
}
